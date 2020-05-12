package usecase

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
	"log"
)

// Service encapsulates user usecase

type AuthUseCase struct {
	userRepo auth.UserRepository
}

// NewService creates a new user usecase
func NewService(userRepo auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo}
}

func (auth *AuthUseCase) Register(userData *model.User) (err error) {
	err = auth.userRepo.Create(userData)
	if err != nil {
		log.Printf("usecase %s", err)
		return
	}

	return nil
}

func (auth *AuthUseCase) FetchUserByEmail(email string) (res *model.User, err error) {
	res, err = auth.userRepo.GetByEmail(email)
	if err != nil {
		return &model.User{}, err
	}
	return
}
