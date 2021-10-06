package users_test

import (
	"daily-tracker-calories/app/middleware/auth"
	"daily-tracker-calories/app/presenter/users"
	usersDom "daily-tracker-calories/bussiness/users"
	_mocksUser "daily-tracker-calories/bussiness/users/mocks"
	"daily-tracker-calories/helper"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	_ "time"
)

var (
	mockUserService  _mocksUser.Service
	userHandler      *users.Presenter
	jwtAuth          auth.ConfigJWT
	validate         validator.Validate
	userRequest      string
	userLoginRequest string
	userDomain       usersDom.Domain
)

func TestMain(m *testing.M) {
	userHandler = users.NewHandler(&mockUserService, &jwtAuth, &validate)
	userRequest = `{
	 "Name": "testuser",
    "Email": "test@mail.com",
    "Password": "test",
    "Gender": "male"
	}`
	userLoginRequest = `{
	 "Email": "test@mail.com",
    "Password": "test"
	}`
	hashPass, _ := helper.PasswordHash("test")
	userDomain = usersDom.Domain{
		ID:       1,
		Name:     "testuser",
		Email:    "test@mail.com",
		Password: hashPass,
		Avatar:   "images/avatar/ava.jpg",
		Gender:   "male",
	}
	m.Run()
}
func TestRegister(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users/register", strings.NewReader(userRequest))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e := echo.New()
		c := e.NewContext(req, rec)
		mockUserService.On("RegisterUser", mock.Anything, mock.Anything).Return(&userDomain, nil).Once()
		data := helper.Meta{
			Message: "Success Register User",
			Code:    http.StatusOK,
			Status:  "Success",
		}
		resp := helper.Response{
			Meta: data,
			Data: &userDomain,
		}
		expected, _ := json.Marshal(resp)
		if assert.NoError(t, userHandler.RegisterUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string(expected), rec.Body.String())
		}
	})
}
