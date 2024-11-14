package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Completed   bool `gorm:"default:false"`
}

func NewTask(title, description string) *Task {
	return &Task{
		Title:       title,
		Description: description,
	}
}
