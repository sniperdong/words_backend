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
				m.PUT("/v", newsBiz.UpVideo)
				m.GET("/v/like", newsBiz.LikeVideo)
				m.GET("/v/star", newsBiz.StarVideo)
				m.POST("/v/log", newsBiz.AddVideoLog)
				m.GET("/v/log", newsBiz.GetVideoLog)
				m.GET("/v/log/slice", newsBiz.GetVideoLogSlice)
				m.GET("/v/log/like", newsBiz.LikeVideoReply)
				m.GET("/v/log/star", newsBiz.StarVideoReply)
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
