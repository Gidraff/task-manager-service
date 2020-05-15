package usecase

import (
	"github.com/Gidraff/task-manager-service/auth/repository/mock"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUseCase_SignUp(t *testing.T) {
	repo := new(mock.UserRepoMock)
	uc := NewUseCase(repo)

	var (
		username = "john"
		email    = "johndoe@gmail.com"
		password = "qw123d4rdt45kfj2gw4rt"

		user = &model.User{
			ID:        0,
			Username:  username,
			Email:     email,
			Password:  password,
			CreatedAt: time.Time{},
		}
	)

	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(user)
	assert.NoError(t, err)

	repo.On("CreateUser", user).Return()

}
