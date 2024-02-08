package store

import (
	"echo.osaxon/model"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db}
}

func (us *UserStore) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := us.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
