package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

// UseCase type
type UseCase interface {
	Register(username, email, password string) error
	GetUserByEmail(email string) (model.User, error)
}
