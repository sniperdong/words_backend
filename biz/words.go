package biz

import (
	"context"
	"words_backend/dao"
	"words_backend/dao/model"
	"words_backend/util"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetWords(ctx context.Context, c *app.RequestContext) {
	words, err := dao.GetWords(ctx)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	util.SuccessResponse(ctx, c, words)
}
func AddWords(ctx context.Context, c *app.RequestContext) {
	var req AddWordsRequest
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	id, err := dao.AddWord(ctx, &model.Word{
		Word:              req.Word,
		Means:             req.Means,
		Pronounce:         req.Pronounce,
		Property:          req.Property,
		Sentences:         req.Sentences,
		Plural:            req.Plural,
		PastTense:         req.PastTense,
		PastParticiple:    req.PastParticiple,
		PresentParticiple: req.PresentParticiple,
	})
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	util.SuccessResponse(ctx, c, id)
}
