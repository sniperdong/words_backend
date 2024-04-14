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

func UpPublish(ctx context.Context, id uint, publish int) error {
	res := newsDB.Model(&model.Videos{ID: id}).Updates(map[string]interface{}{"publish": publish})
	affect, err := res.RowsAffected, res.Error
	if err != nil {
		return err
	}

	if affect == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func UpMemo(ctx context.Context, id uint, memo string) error {
	res := newsDB.Model(&model.Videos{ID: id}).Updates(map[string]interface{}{"memo": memo})
	affect, err := res.RowsAffected, res.Error
	if err != nil {
		return err
	}

	if affect == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func UpContent(ctx context.Context, id uint, content string) error {
	res := newsDB.Model(&model.Videos{ID: id}).Updates(map[string]interface{}{"content": content})
	affect, err := res.RowsAffected, res.Error
	if err != nil {
		return err
	}

	if affect == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
