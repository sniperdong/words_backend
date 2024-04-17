package news

import (
	"context"
	"errors"
	"fmt"
	"time"
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
	Likes   int    `json:"likes"`
	Stars   int    `json:"stars"`
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
			Likes:   videos[i].Likes,
			Stars:   videos[i].Stars,
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

type UpVideoReq struct {
	Id      uint    `json:"id"`
	Publish *bool   `json:"publish"`
	Memo    *string `json:"memo"`
	Content *string `json:"content"`
	Name    *string `json:"name"`
}

func UpVideo(ctx context.Context, c *app.RequestContext) {
	var req UpVideoReq
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	var publish *int
	if req.Publish != nil {
		reqPublish := dao.PublishVideoYes
		if !*req.Publish {
			reqPublish = dao.PublishVideoNo
		}
		publish = &reqPublish
	}
	err := dao.UpVideo(ctx, req.Id, publish, req.Name, req.Memo, req.Content)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	util.SuccessResponse(ctx, c, nil)
}

type AddVideoLogReq struct {
	VideoId uint   `json:"id"`
	Content string `json:"content"`
}

func AddVideoLog(ctx context.Context, c *app.RequestContext) {
	var req AddVideoLogReq
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	if len(req.Content) == 0 {
		util.FailResponse(ctx, c, errors.New("empty content"))
		return
	}
	if _, err := dao.AddVideosLog(ctx, &model.VideoLogs{VideoID: req.VideoId, Content: req.Content}); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	util.SuccessResponse(ctx, c, nil)
}

type GetVideoLogReq struct {
	VideoId uint `json:"id" query:"id"`
}
type VideoLog struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Likes     int       `json:"likes"`
	Stars     int       `json:"stars"`
}
type GetVideoLogRes struct {
	Video *VideoDetail `json:"video"`
	Logs  []*VideoLog  `json:"logs"`
}

func GetVideoLog(ctx context.Context, c *app.RequestContext) {
	var req GetVideoLogReq
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	video, err := dao.GetVideo(ctx, req.VideoId)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	var res GetVideoLogRes
	res.Video = &VideoDetail{
		ID:      video.ID,
		Name:    video.Name,
		Memo:    video.Memo,
		Content: video.Content,
		Path:    video.Path,
		Likes:   video.Likes,
		Stars:   video.Stars,
	}
	logs, err := dao.GetVideosLogs(ctx, req.VideoId)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	res.Logs = make([]*VideoLog, len(logs))
	for i := range logs {
		res.Logs[i] = &VideoLog{
			ID:        logs[i].ID,
			Content:   logs[i].Content,
			CreatedAt: logs[i].CreatedAt,
			UpdatedAt: logs[i].UpdatedAt,
			Likes:     logs[i].Likes,
			Stars:     logs[i].Stars,
		}
	}
	util.SuccessResponse(ctx, c, res)
}

func GetVideoLogSlice(ctx context.Context, c *app.RequestContext) {
	var req struct {
		VideoID   uint `query:"video_id"`
		LastLogID uint `query:"last_log_id"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}

	logs, err := dao.GetVideosNewLogs(ctx, req.VideoID, req.LastLogID)
	if err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	res := make([]*VideoLog, len(logs))
	for i := range logs {
		res[i] = &VideoLog{
			ID:        logs[i].ID,
			Content:   logs[i].Content,
			CreatedAt: logs[i].CreatedAt,
			UpdatedAt: logs[i].UpdatedAt,
			Likes:     logs[i].Likes,
			Stars:     logs[i].Stars,
		}
	}
	util.SuccessResponse(ctx, c, res)
}

func LikeVideo(ctx context.Context, c *app.RequestContext) {
	var req struct {
		ID uint `query:"id"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	dao.AddVideoLikes(ctx, req.ID)
	util.SuccessResponse(ctx, c, nil)
}
func StarVideo(ctx context.Context, c *app.RequestContext) {
	var req struct {
		ID uint `query:"id"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	dao.AddVideoStars(ctx, req.ID)
	util.SuccessResponse(ctx, c, nil)
}

func LikeVideoReply(ctx context.Context, c *app.RequestContext) {
	var req struct {
		ID uint `query:"id"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	dao.AddVideoLogsLikes(ctx, req.ID)
	util.SuccessResponse(ctx, c, nil)
}
func StarVideoReply(ctx context.Context, c *app.RequestContext) {
	var req struct {
		ID uint `query:"id"`
	}
	if err := c.BindAndValidate(&req); err != nil {
		util.FailResponse(ctx, c, err)
		return
	}
	dao.AddVideoLosgStars(ctx, req.ID)
	util.SuccessResponse(ctx, c, nil)
}
