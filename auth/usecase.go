package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

type BasicAuthUseCase interface {
	SignUp(userData *model.User) error
	FetchUserByEmail(email string) (res *model.User, err error)
}
