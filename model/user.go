package model

import "time"

// User model
type Account struct {
	ID        uint
	Username  string `json:"username" validate:"min=3,max=50,required"`
	Email     string `json:"email" validate:"email,required"`
	Password  string `json:"password" validate:"min=8,required"`
	Active    bool   `json:"active" validate:"required"`
	CreatedOn time.Time
	LastLogin time.Time
}
