package repository

import (
	"database/sql"
	"time"
	"time-tracker/internal/model"
)

type Users interface {
	GetUsers(filters map[string]string, limit int, cursor int) ([]model.User, error)
	GetUserById(userId int) (model.User, error)
	CreateUser(user model.User) (userId int, err error)
	DeleteUser(userId int) error
	UpdateUser(userId int, user model.UpdateUserInput) error
}

type Tasks interface {
	GetTasksForPeriod(userId int, start, end time.Time) ([]model.Task, error)
	CreateTask(task model.Task) (model.Task, error)
	FinishTask(userId int, taskId int) (model.Task, error)
}

type Repository struct {
	Users
	Tasks
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
		Tasks: NewTasksRepo(db),
	}
}
