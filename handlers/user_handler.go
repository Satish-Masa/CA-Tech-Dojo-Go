package handlers

import (
	"CA-Tech-Dojo-Go/config"
	"CA-Tech-Dojo-Go/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func creatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req models.UserCreatRequest

		row, err := ioutil.ReadFile("./user_creat_reeequest.json")
		if err != nil {
			fmt.Println(err.Error())
		}

		if err := json.Unmarshal(row, &req); err != nil {
			fmt.Println(err.Error())
		}

		token, err := models.CreatToken(string(row))
		if err != nil {
			fmt.Println(err.Error())
		}

		var resp models.UserCreatResponse

		resp.Token = token

		u := models.NewUser(string(row), token)
		db := ConnectDB()
		defer db.Close()
		db.Create(&u)

		v, err := json.Marshal(resp)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(v)
	} else {
		fmt.Println("not Method")
	}
}

func ConnectDB() *gorm.DB {
	tmp := "%s:%s@%s/%s"
	connect := fmt.Sprintf(tmp, config.Config.DbUser, config.Config.Password, config.Config.Tcp, config.Config.DbName)
	driver := config.Config.SQLDriver
	db, _ := gorm.Open(driver, connect)
	db.AutoMigrate(&models.User{})
	return db
}

func Router() {
	http.HandleFunc("/user/creat", creatHandler)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil))
}
