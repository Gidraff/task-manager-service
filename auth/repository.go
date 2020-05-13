package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

// AuthRepository encapsulates the logic to access user from the data source
type BasicAuthRepository interface {
	CreateUser(u *model.User) error
	GetUserByEmail(email string) (res *model.User, err error)
}
