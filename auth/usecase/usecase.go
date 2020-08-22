package usecase

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
)

// encapsulates user usecase
type UseCase struct {
	userAuthRepo auth.UserRepository
}

// NewService creates a new user usecase
func NewUseCase(userAuthRepo auth.UserRepository) *UseCase {
	return &UseCase{userAuthRepo}
}

func (authUC *UseCase) RegisterAccount(username, email, password string) error {
	err := authUC.userAuthRepo.Store(username, email, password)
	if err != nil {
		return err
	}
	return nil
}

func (authUC *UseCase) GetAccountByEmail(email string) (*model.Account, error) {
	user, err := authUC.userAuthRepo.FetchByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
