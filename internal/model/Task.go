package model

import "time"

type Task struct {
	Id        int       `json:"id" `
	Name      string    `json:"name" `
	UserId    int       `json:"user_id" `
	StartTime time.Time `json:"start_time" `
	EndTime   time.Time `json:"end_time" `
}

type TaskSummary struct {
	TaskName string `json:"taskName" `
	Duration string `json:"duration" example:"04:20"`
	Start    string `json:"start" example:"2024-07-11 13:10"`
	End      string `json:"end" example:"2024-07-11 17:30"`
}

type TaskDataToCreate struct {
	Name string `json:"name"`
}
