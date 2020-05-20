package server

import (
	"fmt"
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func creatHandler(c echo.Context) error {
	req := new(models.UserCreatRequest)

	if err := c.Bind(req); err != nil {
		return err
	}

	token, err := models.CreatToken(req.name)
	if err != nil {
		return err
	}

	resp := new(models.UserCreatResponse)

	resp.Token = token

	u := models.NewUser(req.name, token)
	db := ConnectDB()
	defer db.Close()
	db.Create(&u)

	if err := c.Bind(resp); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, resp)
}

func ConnectDB() *gorm.DB {
	tmp := "%s:%s@%s/%s"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.Tcp, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, _ := gorm.Open(driver, connect)
	db.AutoMigrate(&models.User{})
	return db
}

func getHandler(c echo.Context) error {
	u := new(models.User)

	db := ConnectDB()
	defer db.Close()

	if err := c.Bind(u); err != nil {
		return err
	}
	db.First(&u, "name = ?", u.Token)

	resp := new(models.UserGetResponce)
	resp.Name = u.Name

	return c.JSON(http.StatusOK, resp)
}

func updateHandler(c echo.Context) error {
	u := new(models.User)

	db := ConnectDB()
	defer db.Close()

	if err := c.Bind(u); err != nil {
		return err
	}
	db.First(&u, "name = ?", u.Token)

	req := new(models.UserUpdateRequest)

	if err := c.Bind(req); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, req)
}

func Start() {
	e := echo.New()
	e.POST("/user/creat", creatHandler)
	e.GET("/user/get", getHandler)
	e.PUT("/user/update", updateHandler)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
