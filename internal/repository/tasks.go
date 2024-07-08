package repository

import (
	"time-tracker/internal/model"

	"gorm.io/gorm"
)

type TasksRepo struct {
	db *gorm.DB
}

func NewTasksRepo(db *gorm.DB) *TasksRepo {
	return &TasksRepo{db: db}
}

func (r *TasksRepo) CreateTask(task model.Task) (model.Task, error) {

	result := r.db.Create(&task)
	if result.Error != nil {
		return model.Task{}, result.Error
	}

	return task, nil

}
