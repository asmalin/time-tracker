package service

import "time-tracker/internal/repository"

type TasksService struct {
	repo repository.Tasks
}

func NewTasksService(repo repository.Tasks) *TasksService {
	return &TasksService{repo: repo}
}
