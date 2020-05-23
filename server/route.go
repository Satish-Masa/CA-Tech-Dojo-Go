package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/models"
	"github.com/labstack/echo/v4"
)

func creatHandler(c echo.Context) error {
	req := new(models.UserCreatRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	resp := models.FetchToken(req.Name)

	err := models.SaveUser(req.Name, resp.Token)
	if err != nil {
		log.Println(err)
	}

	if err := c.Bind(resp); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resp)
}

func getHandler(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	resp := models.SearchUser(u.Token)

	if err := c.Bind(resp); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func updateHandler(c echo.Context) error {
	req := new(models.UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	return models.UpdateUser(req.Name, req.Token)
}

func Start() {
	e := echo.New()
	e.POST("/user/creat", creatHandler)
	e.GET("/user/get", getHandler)
	e.PUT("/user/update", updateHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
