package usecase

import (
	"github.com/Gidraff/task-manager-service/auth"
	"github.com/Gidraff/task-manager-service/model"
)

// S encapsulates user usecase
type UseCase struct {
	userAuthRepo auth.UserRepository
}

// NewService creates a new user usecase
func NewUseCase(userAuthRepo auth.UserRepository) *UseCase {
	return &UseCase{userAuthRepo}
}

func (authUC *UseCase) SignUp(user *model.User) (err error) {
	err = authUC.userAuthRepo.CreateUser(user)
	if err != nil {
		return
	}

	return nil
}

//func (auth *AuthUseCase) FetchUserByEmail(email string) (res *model.User, err error) {
//	res, err = auth.authRepo.GetUserByEmail(email)
//	if err != nil {
//		return &model.User{}, err
//	}
//	return
//}
