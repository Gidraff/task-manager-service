package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

type UseCase interface {
	Register(userData *model.User) error
	FetchUserByEmail(email string) (res *model.User, err error)
}
