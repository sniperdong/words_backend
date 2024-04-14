package news

import (
	"context"
	"fmt"
	"words_backend/dao"
	"words_backend/dao/model"
	"words_backend/util"

	"github.com/cloudwego/hertz/pkg/app"
)

type GetAllRequest struct {
	KeyWord  string `json:"key" query:"key"`
	Page     int    `json:"page" query:"page"`
	PageSize int    `json:"page_size" query:"page_size"`
}

type VideoDetail struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Memo    string `json:"memo"`
	Content string `json:"content"`
	Path    string `json:"path"`
	Publish bool   `json:"publish"`
}

func GetAll(ctx context.Context, c *app.RequestContext) {
	var req GetAllRequest
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	videos, err := dao.GetVideos(ctx, req.KeyWord, dao.PublishVideoIgnore, req.Page, req.PageSize)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	res := make([]*VideoDetail, len(videos))
	for i := range videos {
		res[i] = &VideoDetail{
			ID:      videos[i].ID,
			Name:    videos[i].Name,
			Memo:    videos[i].Memo,
			Content: videos[i].Content,
			Path:    videos[i].Path,
			Publish: videos[i].Publish == dao.PublishVideoYes,
		}
	}
	util.SuccessResponse(ctx, c, res)
}

type TotalRequest struct {
	KeyWord string `json:"key" query:"key"`
}

func Total(ctx context.Context, c *app.RequestContext) {
	var req TotalRequest
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	total, err := dao.TotalVideos(ctx, req.KeyWord, dao.PublishVideoIgnore)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	util.SuccessResponse(ctx, c, total)
}
func AddVideo(ctx context.Context, c *app.RequestContext) {
	// single file
	file, _ := c.FormFile("file")

	// Upload the file to specific dst
	if err := c.SaveUploadedFile(file, fmt.Sprintf("./static/%s", file.Filename)); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	id, err := dao.AddVideo(ctx, &model.Videos{
		Name: file.Filename,
		Path: fmt.Sprintf("/static/%s", file.Filename),
	})
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	util.SuccessResponse(ctx, c, id)
}

func BatchAddVideo(ctx context.Context, c *app.RequestContext) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		fmt.Println(file.Filename)

		// Upload the file to specific dst.
		if err := c.SaveUploadedFile(file, fmt.Sprintf("./static/%s", file.Filename)); err != nil {
			util.FailResponse(ctx, c, err)
			return
		}

		_, err := dao.AddVideo(ctx, &model.Videos{
			Name: file.Filename,
			Path: fmt.Sprintf("/static/%s", file.Filename),
		})
		if err != nil {
			util.FailResponse(ctx, c, err)
			return
		}
	}

	util.SuccessResponse(ctx, c, nil)
}

type UpVideo struct {
	Id      uint    `json:"id"`
	Publish *bool   `json:"publish"`
	Memo    *string `json:"memo"`
	Content *string `json:"content"`
}

func UpPublish(ctx context.Context, c *app.RequestContext) {
	var req UpVideo
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	if req.Publish == nil {
		util.FailResponse(ctx, c, fmt.Errorf("publish of %d is not found", req.Id))
		return
	}

	publish := dao.PublishVideoYes
	if !*req.Publish {
		publish = dao.PublishVideoNo
	}
	err := dao.UpPublish(ctx, req.Id, publish)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	util.SuccessResponse(ctx, c, nil)
}
func UpMemo(ctx context.Context, c *app.RequestContext) {
	var req UpVideo
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	if req.Memo == nil {
		util.FailResponse(ctx, c, fmt.Errorf("memo of %d is not found", req.Id))
		return
	}

	err := dao.UpMemo(ctx, req.Id, *req.Memo)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	util.SuccessResponse(ctx, c, nil)
}
func UpContent(ctx context.Context, c *app.RequestContext) {
	var req UpVideo
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	if req.Content == nil {
		util.FailResponse(ctx, c, fmt.Errorf("content of %d is not found", req.Id))
		return
	}

	err := dao.UpContent(ctx, req.Id, *req.Content)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	util.SuccessResponse(ctx, c, nil)
}
