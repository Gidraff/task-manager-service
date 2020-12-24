package model

import (
	//"github.com/Gidraff/task-manager-service/pkg/utils/helpers"
	"time"
)

// Project encapsulate type project
type Project struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:100;not null;unique" json:"name"`
	Description string    `gorm:"size:255;" json:"description"`
	UserID      uint64    `json:"user_id"`
	Tasks       []Task    `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tasks"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
}
