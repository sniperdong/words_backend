package jwt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"words_backend/dao"
	"words_backend/dao/model"
	"words_backend/util"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
	JWT "github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *JWT.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = JWT.New(&JWT.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			util.SuccessResponse(ctx, c,
				utils.H{
					"code":    code,
					"token":   token,
					"expire":  expire.Format(time.RFC3339),
					"message": "success",
				})
			// c.JSON(http.StatusOK, utils.H{
			// 	"code":    code,
			// 	"token":   token,
			// 	"expire":  expire.Format(time.RFC3339),
			// 	"message": "success",
			// })
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct model.User
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			user, err := dao.FindUsreByName(ctx, loginStruct.Name)
			if err != nil {
				return nil, err
			}
			if user == nil || user.Password != util.MD5(loginStruct.Password) {
				return nil, errors.New("user already exists or wrong password")
			}

			return user, nil
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := JWT.ExtractClaims(ctx, c)
			return &model.User{
				Name: claims[IdentityKey].(string),
			}
		},
		PayloadFunc: func(data interface{}) JWT.MapClaims {
			if v, ok := data.(*model.User); ok {
				return JWT.MapClaims{
					IdentityKey: v.Name,
				}
			}
			return JWT.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt biz err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	var req RegisterRequest
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	if len(req.Name) == 0 || len(req.Password) == 0 {
		util.FailResponse(ctx, c, fmt.Errorf("name and password must be set"))
		return
	}

	if user, _ := dao.FindUsreByName(ctx, req.Name); user != nil {
		util.FailResponse(ctx, c, fmt.Errorf("the user exists"))
		return
	}

	if _, err := dao.AddUser(ctx, &model.User{Name: req.Name, Password: util.MD5(req.Password)}); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	util.SuccessResponse(ctx, c, nil)
}
