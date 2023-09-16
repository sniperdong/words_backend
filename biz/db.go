package biz

import (
	"context"
	"words_backend/dao"
	"words_backend/util"

	"github.com/cloudwego/hertz/pkg/app"
)

func EditDB(ctx context.Context, c *app.RequestContext) {
	var req EditDBRequest
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	if err := dao.DDLExe(ctx, req.SQL); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	util.SuccessResponse(ctx, c, nil)
}

func GetDB(ctx context.Context, c *app.RequestContext) {
	tables, err := dao.GetAllTables(ctx)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	util.SuccessResponse(ctx, c, tables)
}
