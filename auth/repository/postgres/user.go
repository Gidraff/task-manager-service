package postgres

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
	"gorm.io/gorm"
	"time"
)

type userRepo struct {
	Conn *gorm.DB
	//Logger *log.Logger
}

// NewUserRepo returns a new UserRepository interface
func NewUserRepo(db *gorm.DB) auth.UserRepository {
	return &userRepo{Conn: db}
}

// Store persists a new user to db
func (ur *userRepo) Store(username, email, password string) error {
	user := model.User{Username: username, Email: email, Password: password, CreatedAt: time.Now()}
	result := ur.Conn.Debug().Create(&user)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

// FetchByEmail returns a user with the provided email
func (ur *userRepo) FetchByEmail(value string) (model.User, error) {
	user := model.User{}
	result := ur.Conn.Where("email = ? ", value).First(&user)
	err := result.Error
	//err = ur.Conn.QueryRow(`SELECT user_id, username, email, password, active FROM users WHERE email=$1`, value).Scan(&id, &username, &email, &password, &active)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
