package users_test

import (
	"daily-tracker-calories/bussiness/users"
	_userMock "daily-tracker-calories/bussiness/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

var (
	userRepository _userMock.Repository
	userService    users.Service
	domainTest     users.Domain
)

func TestMain(m *testing.M) {
	userService = users.NewService(&userRepository)
	domainTest = users.Domain{
		ID:       1,
		Name:     "budi",
		Email:    "test@mail.com",
		Password: "secret",
		Gender:   "male",
	}
	os.Exit(m.Run())
}

func TestGetByID(t *testing.T) {
	t.Run("Valid Test Find By ID", func(t *testing.T) {
		userRepository.On("FindByID", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		result, err := userService.FindByID(1)
		assert.Nil(t, err)
		assert.Equal(t, domainTest.ID, result.ID)
	})
	t.Run("Valid Test FindByEmail", func(t *testing.T) {
		userRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		result, err := userService.Login("test@mail.com", "secret")
		assert.Nil(t, err)
		assert.Equal(t, domainTest.ID, result.ID)
	})
	t.Run("Valid Register", func(t *testing.T) {
		userRepository.On("Register User", mock.Anything, mock.Anything).Return(&domainTest, nil).Once()
		req := &users.Domain{
			ID:        1,
			Name:      "budi",
			Email:     "test@mail.com",
			Password:  "secret",
			Gender:    "male",
		}
		result, err := userService.RegisterUser(req)
		assert.Nil(t, err)
		assert.Equal(t, domainTest.ID, result)
	})

}
