package interfaces

import (
	"fmt"
	"log"
	"net/http"

	application "github.com/Satish-Masa/CA-Tech-Dojo-Go/application/user"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/domain"
	"github.com/labstack/echo/v4"
)

func creatHandler(c echo.Context) error {
	req := new(application.UserCreatRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	resp, err := application.FetchToken(req)

	if err != nil {
		log.Println(err)
	} else {
		err := saveHandler(c, req, resp)
		return err
	}

}

func saveHandler(c echo.Context, req *application.UserCreatRequest, resp *application.UserCreatResponse) error {
	err := application.SaveUser(req.Name, resp.Token)
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

	resp := application.SearchUser(u)

	if err := c.Bind(resp); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func updateHandler(c echo.Context) error {
	req := new(application.UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	return application.UpdateUser(req.Name, req.Token)
}

func Start() {
	e := echo.New()
	e.POST("/user/creat", creatHandler)
	e.GET("/user/get", getHandler)
	e.PUT("/user/update", updateHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
