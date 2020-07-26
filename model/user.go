package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// gorm.Model
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User model
type User struct {
	gorm.Model
	Username  string    `json:"username" validate:"min=3,max=50,required"`
	Email     string    `json:"email" validate:"email,required"`
	Password  string    `json:"password" validate:"min=8,required"`
	Status bool 		`json:"status"`

}

// TableName returns table name
func (User) TableName() string {
	return "users"
}

func (u *User) Disable() {
	u.Status = false
}

func (u *User) Enable() {
	u.Status = true
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
