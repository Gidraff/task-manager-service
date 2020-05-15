package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

// UserRepository encapsulates the logic to access user from the data source
type UserRepository interface {
	CreateUser(u *model.User) error
	//GetUserByEmail(email string) (res *model.User, err error)
}
