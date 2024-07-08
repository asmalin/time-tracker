package repository

import (
	"time-tracker/internal/model"

	"gorm.io/gorm"
)

type Users interface {
	GetUsers(filters map[string]string, limit int, cursor int) ([]model.User, error)
	GetUserById(userId int) (model.User, error)
	CreateUser(user model.User) (userId int, err error)
	DeleteUser(userId int) error
	UpdateUser(user model.User) error
}

type Tasks interface {
	CreateTask(task model.Task) (model.Task, error)
}

type Repository struct {
	Users
	Tasks
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
		Tasks: NewTasksRepo(db),
	}
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Task{})
}
