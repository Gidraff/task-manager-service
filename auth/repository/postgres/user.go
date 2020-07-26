package postgres

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	Db *gorm.DB
	//Logger *log.Logger
}

// NewAuthRepo returns a new UserRepository interface
func NewUserRepo(db *gorm.DB) auth.UserRepository {
	return &userRepo{Db: db}
}

// CreateUser persists a new user to db
func (ur *userRepo) Store(u *model.User) error {
	if err := ur.Db.Create(&u).Error; err != nil {
		//ur.Logger.Error("Could not save user")
		return err
	}
	return nil
}

// GetUserByEmail returns a user with the provided email
func (ur *userRepo) FetchByEmail(email string) (*model.User, error) {
	var user model.User
	if err := ur.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
