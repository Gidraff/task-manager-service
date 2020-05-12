package model

import (
	"time"
)

// User model
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username" validate:"min=3,max=50,required"`
	Email     string    `json:"email" validate:"email,required"`
	Password  string    `json:"password" validate:"min=8,required"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName returns table name
func (u *User) TableName() string {
	return "user"
}
