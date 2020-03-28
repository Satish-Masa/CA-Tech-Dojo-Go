package controllers

import (
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

func Route() {
	http.HandleFunc("/user/creat", creatHandler)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil))
}
