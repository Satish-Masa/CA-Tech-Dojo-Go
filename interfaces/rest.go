package interfaces

import (
	"fmt"
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/character"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/gacha"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	domainCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/character"
	domainUser "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/user"
	domainUserCharacter "github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/userCharacter"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Rest struct {
	UserRepository  domainUser.UserRepository
	GachaRepository domainUserCharacter.UserCharacterRepository
	CharaRepository domainCharacter.CharacterRepository
}

func (r Rest) createHandler(c echo.Context) error {
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

	charApplication := character.CharacterApplication{
		Repository: r.CharaRepository,
	}

	id := FindUserID(c)

	count, err := charApplication.CountChara()

	resp, err := application.Gacha(req.Times, id, count)
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

func (r Rest) createCharaHandler(c echo.Context) error {
	req := new(character.CharaCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if req.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name",
		}
	}

	chara := domainCharacter.NewCharacter(req.Name)

	application := character.CharacterApplication{
		Repository: r.CharaRepository,
	}

	err := application.CreateChara(chara)
	if err != nil {
		return err
	}

	return nil
}

func (r Rest) Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/user/create", r.createHandler)
	e.POST("/character/create", r.createCharaHandler)

	user := e.Group("/user")
	user.Use(middleware.JWTWithConfig(Config))

	user.GET("/get", r.getHandler)
	user.PUT("/update", r.updateHandler)
	e.POST("/gacha/draw", r.gachaHandler)
	// e.GET("/character/list", listHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
