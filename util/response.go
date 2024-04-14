package util

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

const (
	StatusOK   = "OK"
	StatusFail = "FAIL"
)

func SuccessResponse(ctx context.Context, c *app.RequestContext, data interface{}) {
	if data != nil {
		c.String(consts.StatusOK, JsonEncode(map[string]interface{}{"Status": StatusOK, "Data": data}))
		return
	}
	c.String(consts.StatusOK, JsonEncode(map[string]interface{}{"Status": StatusOK}))
}

func FailResponse(ctx context.Context, c *app.RequestContext, err error) {
	fmt.Println(err)
	c.String(consts.StatusOK, JsonEncode(map[string]interface{}{"Status": StatusFail, "Msg": err.Error()}))
}
