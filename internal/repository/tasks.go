package repository

import (
	"database/sql"
	"fmt"
	"time"
	"time-tracker/internal/model"
)

type TasksRepo struct {
	db *sql.DB
}

func NewTasksRepo(db *sql.DB) *TasksRepo {
	return &TasksRepo{db: db}
}

func (r *TasksRepo) GetTasksForPeriod(userId int, start, end time.Time) ([]model.Task, error) {
	query := `
		SELECT id, name, user_id, start_time, end_time 
		FROM tasks 
		WHERE user_id = $1 
		  AND start_time >= $2 
		  AND end_time <= $3
		ORDER BY EXTRACT(EPOCH FROM end_time - start_time) DESC`
	rows, err := r.db.Query(query, userId, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Id, &task.Name, &task.UserId, &task.StartTime, &task.EndTime); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tasks, nil
}

func (r *TasksRepo) CreateTask(task model.Task) (model.Task, error) {
	query := `
		INSERT INTO tasks (user_id, name, start_time, end_time)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	var id int
	err := r.db.QueryRow(query, task.UserId, task.Name, task.StartTime, task.EndTime).Scan(&id)
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to create task: %w", err)
	}

	task.Id = id
	return task, nil
}

func (r *TasksRepo) FinishTask(userId int, taskId int) (model.Task, error) {
	var task model.Task

	query := `SELECT id, user_id, name, start_time, end_time FROM tasks WHERE id = $1 AND user_id = $2`
	err := r.db.QueryRow(query, taskId, userId).Scan(&task.Id, &task.UserId, &task.Name, &task.StartTime, &task.EndTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Task{}, fmt.Errorf("task not found")
		}
		return model.Task{}, fmt.Errorf("failed to query task: %w", err)
	}

	if !task.EndTime.IsZero() {
		return model.Task{}, fmt.Errorf("the task has already been completed")
	}

	task.EndTime = time.Now()
	updateQuery := `UPDATE tasks SET end_time = $1 WHERE id = $2`
	_, err = r.db.Exec(updateQuery, task.EndTime, task.Id)
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to update task: %w", err)
	}

	return task, nil
}
