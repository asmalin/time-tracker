package repository

import (
	"fmt"
	"time"
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

func (r *TasksRepo) FinishTask(userId int, taskId int) (model.Task, error) {

	var task model.Task
	result := r.db.Where("id = ? AND user_id = ?", taskId, userId).First(&task)

	if result.Error != nil {
		return model.Task{}, fmt.Errorf("task not found")
	}

	if !task.EndTime.IsZero() {
		return model.Task{}, fmt.Errorf("the task has already been completed")
	}

	task.EndTime = time.Now()

	err := r.db.Save(&task).Error

	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
