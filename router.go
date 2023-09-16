package main

import (
	"context"
	"words_backend/biz"

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
		admin.GET("", biz.GetWords)

		words := api.Group("/words")
		words.POST("", biz.AddWords)
		words.GET("", biz.GetWords)

		db := api.Group("/db")
		{
			db.GET("", biz.GetDB)
			db.POST("", biz.EditDB)
		}
	}

}
