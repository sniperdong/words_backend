package news

import (
	"context"
	"words_backend/dao"
	"words_backend/util"

	"github.com/cloudwego/hertz/pkg/app"
)

type PublishedListRequest struct {
	KeyWord  string `json:"key" query:"key"`
	Page     int    `json:"page" query:"page"`
	PageSize int    `json:"page_size" query:"page_size"`
}
type VideoSimple struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Memo  string `json:"memo"`
	URL   string `json:"url"`
	Likes int    `json:"likes"`
	Stars int    `json:"stars"`
}

func PublishedList(ctx context.Context, c *app.RequestContext) {
	var req PublishedListRequest
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	videos, err := dao.GetVideos(ctx, req.KeyWord, dao.PublishVideoYes, req.Page, req.PageSize)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	res := make([]*VideoSimple, len(videos))
	for i := range videos {
		res[i] = &VideoSimple{
			ID:    videos[i].ID,
			Name:  videos[i].Name,
			Memo:  videos[i].Memo,
			URL:   videos[i].Path,
			Likes: videos[i].Likes,
			Stars: videos[i].Stars,
		}
	}
	util.SuccessResponse(ctx, c, res)
}
