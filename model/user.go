package model

import "time"

// User encapsulate type account
type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:100;not null;" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Projects  []Project `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"projects"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
