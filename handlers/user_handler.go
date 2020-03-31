package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Satish-Masa/CA-Tech-Dojo-Go/config"
	"github.com/Satish-Masa/CA-Tech-Dojo-Go/models"
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

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var u models.User

		db := ConnectDB()
		defer db.Close()

		u.Token = r.Header.Get("x-token")
		db.First(&u, "name = ?", u.Token)

		var resp models.UserGetResponce
		resp.Name = u.Name

		v, err := json.Marshal(resp)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(v)

	} else {
		fmt.Println("Not Method")
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var u models.User

		db := ConnectDB()
		defer db.Close()

		u.Token = r.Header.Get("x-token")
		db.First(&u, "name = ?", u.Token)

		var req models.UserUpdateRequest

		v, err := json.Marshal(req)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(v)

	} else {
		fmt.Println("Not Method Type")
	}
}

func Router() {
	http.HandleFunc("/user/creat", creatHandler)
	http.HandleFunc("/user/get", getHandler)
	http.HandleFunc("/user/update", updateHandler)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil))
}
