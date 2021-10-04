package users

import (
	"daily-tracker-calories/app/middleware/auth"
	_request "daily-tracker-calories/app/presenter/users/request"
	_response "daily-tracker-calories/app/presenter/users/response"
	"daily-tracker-calories/bussiness/users"
	"daily-tracker-calories/helper"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Presenter struct {
	serviceUser users.Service
	jwtAuth     *auth.ConfigJWT
}

func NewHandler(userService users.Service, jwtauth *auth.ConfigJWT) *Presenter {
	return &Presenter{
		serviceUser: userService,
		jwtAuth:     jwtauth,
	}
}

func (handler *Presenter) RegisterUser(echoContext echo.Context) error {
	var req _request.User
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Register", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceUser.RegisterUser(domain)
	if err != nil {
		response := helper.APIResponse("Failed Register", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helper.APIResponse("Success Register User", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) LoginUser(echoContext echo.Context) error {
	var req _request.UserLogin
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp, err := handler.serviceUser.Login(req.Email, req.Password)
	if err != nil {
		response := helper.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := helper.APIResponse("Generate Token Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success Login", http.StatusOK, "Success", _response.UserLogin{Token: resp})
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) UpdateUser(echoContext echo.Context) error {
	var req _request.User
	if err := echoContext.Bind(&req); err != nil {
		response := helper.APIResponse("Failed FindByEmail", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	user := auth.GetUser(echoContext) // ID Get From JWT
	userID := user.ID
	resp, err := handler.serviceUser.Update(userID, domain)
	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helper.APIResponse("Failed Find By ID", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp, err := handler.serviceUser.FindByID(id)
	if err != nil {
		response := helper.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) UploadAvatar(echoContext echo.Context) error {
	file, err := echoContext.FormFile("avatar")
	if err != nil {
		response := helper.APIResponse("Uploud Avatar Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	user := auth.GetUser(echoContext) // ID Get From JWT
	userID := user.ID
	path := fmt.Sprintf("images/avatar/%d-%s", userID, file.Filename)
	src, err := file.Open()
	if err != nil {
		response := helper.APIResponse("Uploud Avatar Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	defer src.Close()
	destination, err := os.Create(path)
	if err != nil {
		response := helper.APIResponse("Uploud Avatar Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	defer destination.Close()
	if _, err = io.Copy(destination, src); err != nil {
		response := helper.APIResponse("Uploud Avatar Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	_, err = handler.serviceUser.UploadAvatar(userID, path)
	if err != nil {
		response := helper.APIResponse("Uploud Avatar Failed", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success", http.StatusOK, "Success", fmt.Sprintf("File %s uploaded successfully with fields User ID : %d and path : %s", file.Filename, user.ID, path))
	return echoContext.JSON(http.StatusOK, response)
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := viper.GetString(`jwt.token`)
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
