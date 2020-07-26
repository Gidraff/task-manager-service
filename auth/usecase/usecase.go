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

func (authUC *UseCase) SignUp(user *model.User) error {
	//authUC.logger.Info("Registering user")
	err := authUC.userAuthRepo.Store(user)
	if err != nil {
		return err
	}
	return nil
}

func (authUC *UseCase) GetUserByEmail(email string) (*model.User, error) {
	//authUC.logger.Info("Get User by email")
	user, err := authUC.userAuthRepo.FetchByEmail(email)
	if err != nil {
		//authUC.logger.Error("Could not find user with email")
		return nil, err
	}
	return user, nil
}
