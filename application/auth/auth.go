package auth

import (
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/interfaces"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	Name string `json: "name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

func creatToken(u *interfaces.UserCreatRequest) (string, error) {
	if u.Name == "" {
		return "", &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name",
		}
	}

	claims := &jwtCustomClaims{
		Name: u.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func FetchToken(u *interfaces.UserCreatRequest) (token string, err error) {
	token, err = creatToken(u)
	if err != nil {
		return "", err
	}

	return token, nil
}

func FindToken(c echo.Context) (string, error) {
	t := c.Get("user").(jwt.Token)
	token, err := t.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
