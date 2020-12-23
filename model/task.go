package model

import "time"

// Task encapsulates type task
type Task struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	ProjectID   uint64    `json:"project_id"`
	UserID      uint64    `json:"user_id"`
	Title       string    `gorm:"size:100;not null;unique" json:"title"`
	Description string    `gorm:"size:255;" json:"description"`
	Priority    string    `gorm:"size:15;default:primary" json:"priority"`
	Status      string    `gorm:"size:15;default:inProgress" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
