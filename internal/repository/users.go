package repository

import (
	"fmt"
	"time-tracker/internal/model"

	"gorm.io/gorm"
)

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) GetUsers(filters map[string]string, limit int, cursor int) ([]model.User, error) {

	var users []model.User

	query := r.db.Model(&model.User{})

	for field, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	query = query.Where("id > ?", cursor)

	if limit != 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
