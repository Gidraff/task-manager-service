package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

type UseCase interface {
	SignUp(u *model.User)  error
	GetUserByEmail(email string) (*model.User, error)
}
