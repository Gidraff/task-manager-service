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

		user = &model.User{
			Username:  username,
			Email:     email,
			Password:  password,
		}
	)

	repo.On("Store", mc.AnythingOfType("*model.User")).Return(nil)
	err := uc.SignUp(user)
	assert.NoError(t, err)
}
