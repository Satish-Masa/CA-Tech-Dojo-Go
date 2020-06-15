package interfaces

import (
	"fmt"
	"net/http"

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

func (r Rest) creatHandler(c echo.Context) error {
	req := new(user.UserCreatRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if req.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name",
		}
	}

	u := domainUser.NewUser(req.Name)

	application := user.UserApplication{
		Repository: r.UserRepository,
	}

	err := application.SaveUser(u)
	if err != nil {
		return err
	}

	resp, err := FetchToken(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, resp)
}

func (r Rest) getHandler(c echo.Context) error {
	uid := FindUserID(c)

	application := user.UserApplication{
		Repository: r.UserRepository,
	}

	resp, err := application.FindUser(uid)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) updateHandler(c echo.Context) error {
	req := new(user.UserUpdateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	uid := FindUserID(c)

	application := user.UserApplication{
		Repository: r.UserRepository,
	}

	return application.UpdateUser(req.Name, uid)
}

func (r Rest) gachaHandler(c echo.Context) error {
	req := new(gacha.GachaDrawRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	application := gacha.GachaApplication{
		Repository: r.GachaRepository,
	}

	id := FindUserID(c)

	resp, err := application.Gacha(req.Times, id)
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
