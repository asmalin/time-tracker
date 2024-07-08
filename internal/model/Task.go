package model

import "time"

type Task struct {
	Id        int       `json:"id" `
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	UserId    int       `json:"user_id"`
	StartTime time.Time `json:"start_time" gorm:"type: timestamp with time zone"`
	EndTime   time.Time `json:"end_time" gorm:"type: timestamp with time zone"`
	User      User      `json:"-" gorm:"foreignKey:UserId;"`
}
