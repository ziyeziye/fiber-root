package model

import "time"

type Task struct {
	ID         uint      `json:"id" gorm:"primaryKey;column:id;type:int unsigned auto_increment"`
	Task       string    `json:"task" gorm:"column:task;type:varchar(255);not null"`
	Status     bool      `json:"status" gorm:"column:status;type:tinyint unsigned;not null;default:0;comment:0,1"` // 0,1
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;type:datetime"`
}

// TableName returns the table name of the Task model
func (t *Task) TableName() string {
	return "task"
}
