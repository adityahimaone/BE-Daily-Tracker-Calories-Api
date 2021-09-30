package auth

import (
	"daily-tracker-calories/app/presenter"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return presenter.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

// GenerateToken jwt ...
func (jwtConf *ConfigJWT) GenerateToken(userID int) string {
	claims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token
}

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

/*func CreateToken(userID int) (token string, err error) {
  	claims := JwtCustomClaims{
  		userID,
  		jwt.StandardClaims{
  			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(viper.GetInt(`jwt.expired`))).Unix(),
  		},
  	}

  	initToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  	token, err = initToken.SignedString([]byte(viper.GetString(`jwt.token`)))
  	return
  }



  func GetKey(token *jwt.Token) int {
  	keyID, _ := token.Header["id"].(int)
  	return keyID
  }
*/
