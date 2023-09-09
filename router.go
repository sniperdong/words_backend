package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterRoute(h *server.Hertz) {
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello words!")
	})

	// h.StaticFS("/", &app.FS{Root: "./", GenerateIndexPages: true})
	api := h.Group("/api")
	{
		admin := api.Group("/admin")
		admin.GET("", func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "admin")
		})

		words := api.Group("/words")
		words.GET("", func(c context.Context, ctx *app.RequestContext) {
			ctx.String(consts.StatusOK, "words")
		})
	}

}
