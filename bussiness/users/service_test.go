package users_test

import (
	"daily-tracker-calories/app/middleware/auth"
	"daily-tracker-calories/bussiness/users"
	_ "daily-tracker-calories/bussiness/users"
	_userMocks "daily-tracker-calories/bussiness/users/mocks"
	"daily-tracker-calories/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockUserRepository _userMocks.Repository
	mockUserService    _userMocks.Service
	userService        users.Service
	userDomain         users.Domain
)

func TestMain(m *testing.M) {
	userService = users.NewService(&mockUserRepository, &auth.ConfigJWT{})
	hashPass, _ := helper.PasswordHash("testpassword")
	userDomain = users.Domain{
		ID:       1,
		Name:     "test name",
		Email:    "testing@mail.com",
		Password: hashPass,
		Avatar:   "/images/avatar/test.jpg",
		Gender:   "Male",
	}
	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		mockUserRepository.On("Insert", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		inputUser := users.Domain{
			ID:       1,
			Name:     "test name",
			Email:    "testing@mail.com",
			Password: "testpassword",
			Avatar:   "/images/avatar/test.jpg",
			Gender:   "Male",
		}
		result, _ := userService.RegisterUser(&inputUser)
		assert.Equal(t, &users.Domain{}, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		mockUserRepository.On("Insert", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		inputUser := users.Domain{
			ID:       1,
			Name:     "test name",
			Email:    "testing@mail.com",
			Password: "testpassword",
			Avatar:   "/images/avatar/test.jpg",
			Gender:   "Male",
		}
		result, _ := userService.RegisterUser(&inputUser)
		assert.Equal(t, &users.Domain{}, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&userDomain, assert.AnError).Once()
		mockUserRepository.On("Insert", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		inputUser := users.Domain{
			ID:       1,
			Name:     "test name",
			Email:    "testing@mail.com",
			Password: "testpassword",
			Avatar:   "/images/avatar/test.jpg",
			Gender:   "Male",
		}
		result, _ := userService.RegisterUser(&inputUser)
		assert.Equal(t, &users.Domain{}, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		mockUserRepository.On("Insert", mock.Anything, mock.Anything).Return(&userDomain, assert.AnError).Once()
		inputUser := users.Domain{
			ID:       1,
			Name:     "test name",
			Email:    "testing@mail.com",
			Password: "testpassword",
			Avatar:   "/images/avatar/test.jpg",
			Gender:   "Male",
		}
		result, _ := userService.RegisterUser(&inputUser)
		assert.Equal(t, &users.Domain{}, result)
	})
}

func TestEditUser(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		inputUser := users.Domain{
			ID:       1,
			Name:     "test name",
			Email:    "testing@mail.com",
			Password: "testpassword",
			Avatar:   "/images/avatar/test.jpg",
			Gender:   "Male",
		}
		result, err := userService.EditUser(1, &inputUser)
		assert.Nil(t, err)
		assert.Equal(t, &userDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(&userDomain, assert.AnError).Once()
		inputUser := users.Domain{
			ID:       1,
			Name:     "test name",
			Email:    "testing@mail.com",
			Password: "testpassword",
			Avatar:   "/images/avatar/test.jpg",
			Gender:   "Male",
		}
		result, _ := userService.EditUser(2, &inputUser)
		assert.Equal(t, &users.Domain{}, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		inputUser := users.Domain{
			Email:    "testing@mail.com",
			Password: "testpassword",
		}
		result, err := userService.Login(inputUser.Email, inputUser.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&users.Domain{}, assert.AnError).Once()
		inputUser := users.Domain{
			Email:    "testing@mail.com",
			Password: "testpassword",
		}
		result, err := userService.Login(inputUser.Email, inputUser.Password)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		inputUser := users.Domain{
			Email:    "testing@mail.com",
			Password: "testpassword",
		}
		userDomain.ID = 0
		result, err := userService.Login(inputUser.Email, inputUser.Password)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestEmailAvailable(t *testing.T) {
	t.Run("Valid Test False", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		result, err := userService.EmailAvailable("testing@mail.com")
		assert.Nil(t, err)
		assert.Equal(t, false, result)
	})
	t.Run("Valid Test True", func(t *testing.T) {
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&userDomain, assert.AnError).Once()
		result, _ := userService.EmailAvailable("testing@mail.com")
		assert.Equal(t, true, result)
	})
	t.Run("Valid Test False", func(t *testing.T) {
		resultMock := users.Domain{
			ID:    0,
			Email: "testing@mail.com",
		}
		mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&resultMock, nil).Once()
		result, err := userService.EmailAvailable("testing@mail.com")
		assert.Nil(t, err)
		assert.Equal(t, true, result)
	})
}

func TestUploadAvatar(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		result, err := userService.UploadAvatar(1, userDomain.Avatar)
		assert.Nil(t, err)
		assert.Equal(t, &userDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(&userDomain, assert.AnError).Once()
		result, _ := userService.UploadAvatar(2, userDomain.Avatar)
		assert.Equal(t, &users.Domain{}, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&userDomain, assert.AnError).Once()
		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		result, _ := userService.UploadAvatar(2, userDomain.Avatar)
		assert.Equal(t, &users.Domain{}, result)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("FindByID", mock.Anything).Return(&userDomain, nil).Once()
		result, err := userService.FindByID(1)
		assert.Nil(t, err)
		assert.Equal(t, &userDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("FindByID", mock.Anything).Return(&userDomain, assert.AnError).Once()
		result, _ := userService.FindByID(1)
		assert.Equal(t, &users.Domain{}, result)
	})
}

func TestEditByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything).Return(&userDomain, nil).Once()
		result, err := userService.EditUser(1, &userDomain)
		assert.Nil(t, err)
		assert.Equal(t, &userDomain, result)
	})
}
