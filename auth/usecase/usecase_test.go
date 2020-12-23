package usecase

import (
	"errors"
	"testing"

	"github.com/Gidraff/task-manager-service/auth/repository/mock"
	"github.com/Gidraff/task-manager-service/model"
	"github.com/stretchr/testify/assert"
	mc "github.com/stretchr/testify/mock"
)

func TestUseCase_Register(t *testing.T) {
	repo := new(mock.UserRepoMock)
	uc := NewUseCase(repo)

	user := &model.User{
		Username: "john",
		Email:    "johndoe@gmail.com",
		Password: "qw123d4rdt45kfj2gw4rt",
	}

	t.Run("Success", func(t *testing.T) {
		tempUser := user
		tempUser.ID = 1
		repo.On("FetchByEmail", mc.AnythingOfType("string")).Return(&model.User{}, nil).Once()
		repo.On("Store", mc.Anything, mc.Anything, mc.Anything).Return(nil).Once()

		_, err := uc.GetUserByEmail(user.Email)
		assert.NoError(t, err)
		err = uc.Register(user.Username, user.Email, user.Password)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("duplicate user", func(t *testing.T) {
		existingUser := user
		repo.On("FetchByEmail", mc.AnythingOfType("string")).Return(existingUser, nil).Once()
		repo.On("Store", mc.Anything, mc.Anything, mc.Anything).Return(errors.New("duplicate")).Once()
		err := uc.Register(user.Username, user.Email, user.Password)
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})
}
