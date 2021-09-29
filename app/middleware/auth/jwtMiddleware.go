package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"time"
)


type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func init() {
	viper.SetConfigName("test-config")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func CreateToken(userID int) (token string, err error) {
	claims := jwt.MapClaims{}
	claims["userid"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(viper.GetInt(`jwt.expired`))).Unix()

	initToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = initToken.SignedString([]byte(viper.GetString(`jwt.token`)))
	return
}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
