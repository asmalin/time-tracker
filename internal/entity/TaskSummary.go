package entity

type TaskSummary struct {
	TaskName string `json:"taskName"`
	Duration string `json:"duration"`
	Start    string `json:"start"`
	End      string `json:"end"`
}
