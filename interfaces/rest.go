package interfaces

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/gacha"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/labstack/echo/v4"
)

func creatHandler(c echo.Context) error {
	req := new(user.UserCreatRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	resp, err := user.FetchToken(req)

	if err != nil {
		log.Println(err)
	} else {
		err := saveHandler(c, req, resp)
		return err
	}

}

func saveHandler(c echo.Context, req *user.UserCreatRequest, resp *user.UserCreatResponse) error {
	u, err := domain.NewUser(req.Name, resp.Token)
	if err != nil {
		return err
	}

	err := user.SaveUser(u)
	if err != nil {
		return err
	}

	if err := c.Bind(resp); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resp)
}

func getHandler(c echo.Context) error {
	u := new(domain.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	resp := user.FindUser(u)

	if err := c.Bind(resp); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func updateHandler(c echo.Context) error {
	req := new(user.UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	return user.UpdateUser(req)
}

func gachaHandler(c echo.Context) gacha.GachaDrawResponse {
	req := new(gacha.GachaDrawRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	resp, _ := gacha.DoGacha(req)

	return resp
}

func Start() {
	e := echo.New()
	e.POST("/user/creat", creatHandler)
	e.GET("/user/get", getHandler)
	e.PUT("/user/update", updateHandler)
	e.POST("/gacha/draw", gachaHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
