package rest

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type LoginEndpoint struct {
	JwtKey string
}

func (ep *LoginEndpoint) Login(context echo.Context) error {
	username := context.FormValue("username")
	password := context.FormValue("password")

	if username == "testUser" && password == "1234" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Test User"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(ep.JwtKey))
		if err != nil {
			return err
		}
		return context.JSONPretty(http.StatusOK, map[string]string{
			"token": t,
		}, " ")
	}

	return context.NoContent(http.StatusUnauthorized)
}
