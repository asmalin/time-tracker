package repository

import "gorm.io/gorm"

type TasksRepo struct {
	db *gorm.DB
}

func NewTasksRepo(db *gorm.DB) *TasksRepo {
	return &TasksRepo{db: db}
}

func (r *TasksRepo) GetAllUsers() error {
	return nil
}
