package dao

import (
	"context"
	"words_backend/dao/model"

	"gorm.io/gorm"
)

func AddVideo(ctx context.Context, m *model.Videos) (uint, error) {
	res := newsDB.Create(m)
	if res.Error != nil {
		return 0, res.Error
	}

	return m.ID, nil
}

func GetVideo(ctx context.Context, id uint) (res *model.Videos, err error) {
	if err := newsDB.Where("id=?", id).First(&res).Error; err != nil {
		return nil, err
	}

	return
}

func GetVideos(ctx context.Context, key string, publish int, page, pageSize int) ([]*model.Videos, error) {
	var videos []*model.Videos
	db := newsDB
	if len(key) > 0 {
		db = db.Where("name like ?", "%"+key+"%")
	}
	if publish == 0 || publish == 1 {
		db = db.Where("publish = ?", publish)
	}

	if err := db.Limit(pageSize).Offset(page * pageSize).Find(&videos).Error; err != nil {
		return nil, err
	}

	return videos, nil
}

func TotalVideos(ctx context.Context, key string, publish int) (int64, error) {
	db := newsDB.Model(&model.Videos{})
	if len(key) > 0 {
		db = db.Where("name like ?", "%"+key+"%")
	}
	if publish == 0 || publish == 1 {
		db = db.Where("publish = ?", publish)
	}

	var res int64
	if err := db.Count(&res).Error; err != nil {
		return 0, err
	}

	return res, nil
}

func UpVideo(ctx context.Context, id uint, publish *int, name, memo, content *string) error {
	updates := make(map[string]interface{})
	if publish != nil {
		updates["publish"] = *publish
	}
	if name != nil {
		updates["name"] = *name
	}
	if memo != nil {
		updates["memo"] = *memo
	}
	if content != nil {
		updates["content"] = *content
	}

	if len(updates) == 0 {
		return ErrUpdateNothing
	}

	res := newsDB.Model(&model.Videos{ID: id}).Updates(updates)
	affect, err := res.RowsAffected, res.Error
	if err != nil {
		return err
	}

	if affect == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func AddVideosLog(ctx context.Context, m *model.VideoLogs) (uint, error) {
	res := newsDB.Create(m)
	if res.Error != nil {
		return 0, res.Error
	}

	return m.ID, nil
}

func GetVideosLogs(ctx context.Context, videoID uint) (res []*model.VideoLogs, err error) {
	if err := newsDB.Where("video_id = ?", videoID).Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}

	return
}

func GetVideosNewLogs(ctx context.Context, videoID, LastLogID uint) (res []*model.VideoLogs, err error) {
	if err := newsDB.Where("id > ? and video_id = ?", LastLogID, videoID).Order("id desc").Find(&res).Order("id desc").Error; err != nil {
		return nil, err
	}

	return
}

func AddVideoLikes(ctx context.Context, id uint) error {
	return newsDB.Model(&model.Videos{}).Where("id = ?", id).UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
}

func AddVideoStars(ctx context.Context, id uint) error {
	return newsDB.Model(&model.Videos{}).Where("id = ?", id).UpdateColumn("stars", gorm.Expr("stars + ?", 1)).Error
}

func AddVideoLogsLikes(ctx context.Context, id uint) error {
	return newsDB.Model(&model.VideoLogs{}).Where("id = ?", id).UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
}

func AddVideoLosgStars(ctx context.Context, id uint) error {
	return newsDB.Model(&model.VideoLogs{}).Where("id = ?", id).UpdateColumn("stars", gorm.Expr("stars + ?", 1)).Error
}
