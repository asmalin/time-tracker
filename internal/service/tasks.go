package service

import (
	"time-tracker/internal/model"
	"time-tracker/internal/repository"
)

type TasksService struct {
	repo repository.Tasks
}

func NewTasksService(repo repository.Tasks) *TasksService {
	return &TasksService{repo: repo}
}

func (s *TasksService) StartTask(task model.Task) (model.Task, error) {

	return s.repo.CreateTask(task)
}

func (s *TasksService) FinishTask(userId int, taskId int) (model.Task, error) {
	return s.repo.FinishTask(userId, taskId)
}
