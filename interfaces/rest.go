package interfaces

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/gacha"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain/repository"
	"github.com/labstack/echo/v4"
)

type Rest struct {
	Repository repository.UserRepository
}

func (r Rest) creatHandler(c echo.Context) error {
	req := new(user.UserCreatRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	resp, err := user.FetchToken(req)

	if err != nil {
		log.Println(err)
	} else {
		err := r.saveHandler(c, req, resp)
		return err
	}

}

func (r Rest) saveHandler(c echo.Context, req user.UserCreatRequest, resp user.UserCreatResponse) error {
	u, err := domain.NewUser(req.Name, resp.Token)
	if err != nil {
		return err
	}

	application := user.UserApplication{
		Repository: r.Repository,
	}

	err := application.SaveUser(u)
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
		Repository: r.Repository,
	}

	resp := application.FindUser(u)

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) updateHandler(c echo.Context) error {
	req := new(user.UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	application := user.UserApplication{
		Repository: r.Repository,
	}

	return application.UpdateUser(req)
}

func gachaHandler(c echo.Context) error {
	req := new(gacha.GachaDrawRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	resp, _ := gacha.DoGacha(req)

	return c.JSON(http.StatusOK, resp)
}

func listHandler(c echo.Context) error {
	req := new(domain.User)
	if err := c.Bind(req); err != nil {
		return err
	}

	resp := user.GetList(req)

	return c.JSON(http.StatusOK, resp)
}

func (r Rest) Start() {
	e := echo.New()
	e.POST("/user/creat", r.creatHandler)
	e.GET("/user/get", r.getHandler)
	e.PUT("/user/update", r.updateHandler)
	e.POST("/gacha/draw", gachaHandler)
	e.GET("/character/list", listHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
