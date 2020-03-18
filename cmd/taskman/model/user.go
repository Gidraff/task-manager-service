package model

import (
	"time"
)

// User model
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName returns table name
func (u *User) TableName() string {
	return "user"
}
