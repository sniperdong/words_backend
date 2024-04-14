package main

import (
	"words_backend/biz"
	"words_backend/biz/jwt"
	newsBiz "words_backend/biz/news"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterRoute(h *server.Hertz) {
	api := h.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/login", jwt.JwtMiddleware.LoginHandler)
			user.POST("/register", jwt.Register)
		}

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

		news := api.Group("/news")
		{
			u := news.Group("/u")
			{
				u.GET("/list", newsBiz.PublishedList)
			}
			m := news.Group("/m")
			{
				m.GET("/list", newsBiz.GetAll)
				m.GET("/total", newsBiz.Total)
				m.PUT("/v/up/publish", newsBiz.UpPublish)
				m.PUT("/v/up/content", newsBiz.UpContent)
				m.PUT("/v/up/memo", newsBiz.UpMemo)
			}
		}
	}

	upload := h.Group("/file")
	{
		video := upload.Group("/v")
		{
			video.POST("/single", newsBiz.AddVideo)
			video.POST("/multi", newsBiz.BatchAddVideo)
		}
	}
}
