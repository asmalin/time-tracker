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

func (r *UsersRepo) CreateUser(user model.User) (userId int, err error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.Id, nil
}

func (r *UsersRepo) DeleteUser(userId int) error {

	var user model.User
	result := r.db.First(&user, userId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with ID %d does not exist", userId)
		} else {
			return fmt.Errorf("error: %v", result.Error)
		}
	}

	result = r.db.Delete(model.User{}, "id = ?", userId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UsersRepo) UpdateUser(user model.User) error {
	err := r.db.Save(&user).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepo) GetUserById(userId int) (model.User, error) {
	var user model.User
	result := r.db.First(&user, userId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return model.User{}, fmt.Errorf("user with ID %d does not exist", userId)
		} else {
			return model.User{}, fmt.Errorf("error: %v", result.Error)
		}
	}
	return user, nil
}
