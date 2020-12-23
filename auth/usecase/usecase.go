package usecase

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
)

// UseCase encapsulates user usecase
type UseCase struct {
	userAuthRepo auth.UserRepository
}

// NewUseCase returns a new account usecase
func NewUseCase(userAuthRepo auth.UserRepository) *UseCase {
	return &UseCase{userAuthRepo}
}

// Register register new user
func (authUC *UseCase) Register(username, email, password string) error {
	err := authUC.userAuthRepo.Store(username, email, password)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail returns type User
func (authUC *UseCase) GetUserByEmail(email string) (*model.User, error) {
	account, err := authUC.userAuthRepo.FetchByEmail(email)
	if err != nil {
		return nil, err
	}
	return account, nil
}
