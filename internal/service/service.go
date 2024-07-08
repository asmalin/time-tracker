package service

import (
	"time-tracker/internal/model"
	"time-tracker/internal/repository"
)

type Users interface {
	GetUsers(filters map[string]string, limit int, cursor int) ([]model.User, error)
	GetUserById(userId int) (model.User, error)
	CreateUser(model.User) (userId int, err error)
	DeleteUser(userId int) error
	UpdateUser(userDataToUpdate model.User) error
}

type Tasks interface {
	StartTask(task model.Task) (model.Task, error)
	FinishTask(userId int, taskId int) (model.Task, error)
}

type Service struct {
	Users
	Tasks
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repo.Users),
		Tasks: NewTasksService(repo.Tasks),
	}
}
