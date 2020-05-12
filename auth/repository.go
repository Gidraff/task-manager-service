package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

// UserRepository encapsulates the logic to access user from the data source
type UserRepository interface {
	Create(u *model.User) error
	GetByEmail(email string) (res *model.User, err error)
}
