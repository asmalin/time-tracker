package service

import (
	"fmt"
	"time"
	"time-tracker/internal/entity"
	"time-tracker/internal/model"
	"time-tracker/internal/repository"
)

type TasksService struct {
	repo repository.Tasks
}

func NewTasksService(repo repository.Tasks) *TasksService {
	return &TasksService{repo: repo}
}

func (s *TasksService) GetTasksForPeriod(userId int, start time.Time, end time.Time) ([]entity.TaskSummary, error) {
	tasks := s.repo.GetTasksForPeriod(userId, start, end)

	var taskSummaries []entity.TaskSummary
	for _, task := range tasks {
		if task.StartTime.IsZero() || task.EndTime.IsZero() {
			continue
		}

		duration := task.EndTime.Sub(task.StartTime)
		taskSummaries = append(taskSummaries, entity.TaskSummary{
			TaskName: task.Name,
			Duration: fmt.Sprintf("%02d:%02d", int(duration.Hours()), int(duration.Minutes())%60),
			Start:    task.StartTime.Format("2006-01-02 15:04"),
			End:      task.EndTime.Format("2006-01-02 15:04"),
		})
	}

	if len(taskSummaries) == 0 {
		return nil, fmt.Errorf("no tasks found")
	}

	return taskSummaries, nil
}

func (s *TasksService) StartTask(task model.Task) (model.Task, error) {

	return s.repo.CreateTask(task)
}

func (s *TasksService) FinishTask(userId int, taskId int) (model.Task, error) {
	return s.repo.FinishTask(userId, taskId)
}
