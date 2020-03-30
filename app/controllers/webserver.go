package controllers

import (
	"CA-Tech-Dojo-Go/app/models"
	"CA-Tech-Dojo-Go/config"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func creatHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("app/views/creat.html")
	t.Execute(w, nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UserCreatReqest
	var resp models.UserCreatResponse
	req.Name = r.Form.Get("name")
	resp.Token, _ = models.CreatToken(req.Name)
	models.CreatUser(req.Name, resp.Token)
	http.Redirect(w, r, "/user/get/", http.StatusFound)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Path[len("/user/get/"):]
	var resp models.UserGetResponse
	resp.Name = models.GetUser(token)
	t, _ := template.ParseFiles("app/views/index.html")
	t.Execute(w, resp)
}

func Route() {
	http.HandleFunc("/user/creat", creatHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/user/get", getHandler)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil))
}
