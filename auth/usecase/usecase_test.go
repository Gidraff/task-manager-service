package usecase

import (
	"github.com/Gidraff/task-manager-service/auth/repository/mock"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/assert"
	mc "github.com/stretchr/testify/mock"
	"testing"
)

func TestUseCase_SignUp(t *testing.T) {
	repo := new(mock.UserRepoMock)
	uc := NewUseCase(repo)

	var (
		username = "john"
		email    = "johndoe@gmail.com"
		password = "qw123d4rdt45kfj2gw4rt"

		user = &model.Account{
			Username: username,
			Email:    email,
			Password: password,
		}
	)

	repo.On("Store", mc.Anything, mc.Anything, mc.Anything).Return(nil)
	err := uc.RegisterAccount(user.Username, email, user.Password)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
