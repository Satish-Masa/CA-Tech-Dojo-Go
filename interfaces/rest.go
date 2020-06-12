package interfaces

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/auth"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/gacha"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"
	"github.com/labstack/echo/v4"
)

type Rest struct {
	UserRepository  domainUser.UserRepository
	GachaRepository domainCharacter.CharacterRepository
}

type UserCreatRequest struct {
	Name string `json: "name"`
}

type UserUpdateRequest struct {
	Name  string `json: "name"`
	Token string `json: "token"`
}

type GachaDrawRequest struct {
	Times int `json: "times"`
}

func (r Rest) creatHandler(c echo.Context) error {
	req := new(UserCreatRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if req.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name",
		}
	}

	token, err := auth.FetchToken(req)

	if err != nil {
		log.Println(err)
		return err
	}

	u, err := domain.NewUser(req.Name, token)
	if err != nil {
		return err
	}

	application := user.UserApplication{
		Repository: r.UserRepository,
	}

	resp, err := application.SaveUser(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, resp)
}

func (r Rest) getHandler(c echo.Context) error {
	u := new(domain.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	application := user.UserApplication{
		Repository: r.UserRepository,
	}

	resp := application.FindUser(u)

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) updateHandler(c echo.Context) error {
	req := new(UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	application := user.UserApplication{
		Repository: r.UserRepository,
	}

	return application.UpdateUser(req)
}

func (r Rest) gachaHandler(c echo.Context) error {
	req := new(GachaDrawRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	application := gacha.GachaApplication{
		Repository: r.GachaRepository,
	}

	token, err := auth.FindToken(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to get token",
		}
	}

	resp, err := application.Gacha(req, token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

/* func listHandler(c echo.Context) error {
	req := new(domain.User)
	if err := c.Bind(req); err != nil {
		return err
	}

	resp := user.GetList(req)

	return c.JSON(http.StatusOK, resp)
} */

func (r Rest) Start() {
	e := echo.New()
	e.POST("/user/creat", r.creatHandler)
	e.GET("/user/get", r.getHandler)
	e.PUT("/user/update", r.updateHandler)
	e.POST("/gacha/draw", r.gachaHandler)
	// e.GET("/character/list", listHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
