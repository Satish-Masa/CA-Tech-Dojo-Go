package auth

import (
	"net/http"

	domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserCreatResponse struct {
	Token string `json: "token"`
}

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

func creatToken(u *domainUser.User) (UserCreatResponse, error) {
	if u.Name == "" {
		return UserCreatResponse{}, &echo.HTTPError{
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
		return UserCreatResponse{}, &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to create the token",
		}
	}

	resp := new(UserCreatResponse)
	resp.Token = t

	return *resp, nil
}

func FetchToken(u *domainUser.User) (resp UserCreatResponse, err error) {
	resp, err = creatToken(u)
	if err != nil {
		return UserCreatResponse{}, err
	}

	return resp, nil
}

func FindUserID(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}
