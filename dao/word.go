package dao

import (
	"context"
	"words_backend/dao/model"
)

func GetWords(ctx context.Context) ([]*model.Word, error) {
	var words []*model.Word
	err := db.Find(&words).Error
	if err != nil {
		return nil, err
	}

	return words, nil
}
func AddWord(ctx context.Context, m *model.Word) (uint, error) {
	res := db.Create(m)
	if res.Error != nil {
		return 0, res.Error
	}

	return m.ID, nil
}
