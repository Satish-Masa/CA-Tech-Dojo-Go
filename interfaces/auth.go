package interfaces

import (
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaims struct {
	UID  int    `json: "uid"`
	Name string `json: "name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

func createToken(u *domainUser.User) (user.UserCreatResponse, error) {
	if u.Name == "" {
		return user.UserCreatResponse{}, &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name",
		}
	}

	claims := &jwtCustomClaims{
		UID:  u.ID,
		Name: u.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return user.UserCreatResponse{}, &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to create the token",
		}
	}

	resp := new(user.UserCreatResponse)
	resp.Token = t

	return *resp, nil
}

func FetchToken(u *domainUser.User) (resp user.UserCreatResponse, err error) {
	resp, err = createToken(u)
	if err != nil {
		return user.UserCreatResponse{}, err
	}

	return resp, nil
}

func FindUserID(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}
