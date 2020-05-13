package usecase

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
	"log"
)

// S encapsulates user usecase
type AuthUseCase struct {
	authRepo auth.BasicAuthRepository
}

// NewService creates a new user usecase
func NewAuthUseCase(authRepo auth.BasicAuthRepository) *AuthUseCase {
	return &AuthUseCase{authRepo}
}

func (auth *AuthUseCase) SignUp(userData *model.User) (err error) {
	err = auth.authRepo.CreateUser(userData)
	if err != nil {
		log.Printf("usecase %s", err)
		return
	}

	return nil
}

func (auth *AuthUseCase) FetchUserByEmail(email string) (res *model.User, err error) {
	res, err = auth.authRepo.GetUserByEmail(email)
	if err != nil {
		return &model.User{}, err
	}
	return
}
