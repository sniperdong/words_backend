package dao

import (
	"context"
	"words_backend/dao/model"
)

func AddUser(ctx context.Context, m *model.User) (uint, error) {
	res := db.Create(m)
	if res.Error != nil {
		return 0, res.Error
	}

	return m.ID, nil
}

func FindUsreByName(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	res := db.Model(model.User{Name: name}).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
