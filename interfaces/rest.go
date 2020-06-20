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
	Err "github.com/Satish-Masa/CA-Tech-Dojo-Go/error"
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
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "failed to bind JSON",
		}
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
		switch err {
		case Err.ErrCreateUser:
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to save the user",
			}
		default:
			return err
		}
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
		switch err {
		case Err.ErrFindUser:
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to find the user",
			}
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) updateHandler(c echo.Context) error {
	req := new(user.UserUpdateRequest)
	if err := c.Bind(req); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "failed to bind JSON",
		}
	}

	uid := FindUserID(c)

	application := user.UserApplication{
		Repository: r.UserRepository,
	}

	err := application.UpdateUser(req.Name, uid)
	if err != nil {
		switch err {
		case Err.ErrUpdateUser:
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to update the user",
			}
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"update": "Ok!!",
	})
}

func (r Rest) gachaHandler(c echo.Context) error {
	req := new(gacha.GachaDrawRequest)
	if err := c.Bind(req); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "failed to bind JSON",
		}
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
		switch err {
		case Err.ErrCount:
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to count characters",
			}
		case Err.ErrFindChara:
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to find characters",
			}
		case Err.ErrCreateUserChara:
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to create the userCharacter",
			}
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) listHandler(c echo.Context) error {
	uid := FindUserID(c)

	application := gacha.GachaApplication{
		Repository: r.GachaRepository,
	}

	resp, err := application.FindAll(uid)
	if err != nil {
		switch err {
		case Err.ErrFindAll:
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to find your characters",
			}
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) createCharaHandler(c echo.Context) error {
	req := new(character.CharaCreateRequest)
	if err := c.Bind(req); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "failed to bind JSON",
		}
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
		switch err {
		case Err.ErrCreateChara:
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to save the character",
			}
		default:
			return err
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"Create": "Ok!!",
	})
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
	user.POST("/gacha", r.gachaHandler)
	user.GET("/list", r.listHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
