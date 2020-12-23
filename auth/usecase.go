package auth

import (
	"github.com/Gidraff/task-manager-service/model"
)

type UseCase interface {
	RegisterAccount(username, email, password string) error
	GetAccountByEmail(email string) (*model.Account, error)
}
