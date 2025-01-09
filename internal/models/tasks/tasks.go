package tasks

import (
	"time"
)

// TaskType a task, for example: delete logs
type TaskType struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Kind string `gorm:"default:'task_type'" json:"kind"`
	Name string `gorm:"not null;unique" json:"name"`
}

// Task refers to the specific actions performed by the task, for example: delete logs at 10pm
type Task struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Kind      string    `gorm:"default:'task'" json:"kind"`
	TaskType  string    `gorm:"not null;default:''" json:"taskType"` // related tasks
	Name      string    `gorm:"not null;unique" json:"name"`
	Mode      string    `gorm:"not null;default:'once'" json:"mode"` // once or periodic
	CronSpec  string    `gorm:"not null;default:''" json:"cronSpec"`
	Payload   string    `gorm:"type:text;not null" json:"payload"` // json format
	Result    bool      `gorm:"default:false" json:"result"`
	Event     string    `gorm:"not null;default:''" json:"event"`
	CreatedAt time.Time `gorm:"autoCreateTime;unchangeable" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
}
