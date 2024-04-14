package main

import (
	"context"
	"time"
	"words_backend/biz/jwt"
	"words_backend/dao"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func init() {
	dao.Init()
	jwt.InitJwt()
}
func main() {
	// server.Default() creates a Hertz with recovery middleware.
	// If you need a pure hertz, you can use server.New()
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"), server.WithMaxRequestBodySize(500<<20))

	// static
	Static(h)

	// api
	RegisterRoute(h)

	h.Spin()
}

func Static(h *server.Hertz) {
	// h.Static("/static", "./")

	h.StaticFS("/static", &app.FS{
		Root:        "./static",
		PathRewrite: app.NewPathSlashesStripper(1),
		PathNotFound: func(_ context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusNotFound, "The requested resource does not exist")
		},
		CacheDuration:        time.Second * 5,
		IndexNames:           nil, //[]string{"a.mp4"}
		Compress:             true,
		CompressedFileSuffix: "hertz",
		AcceptByteRange:      true,
	})
}
