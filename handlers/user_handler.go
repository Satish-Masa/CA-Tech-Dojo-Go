package handlers

import (
	"CA-Tech-Dojo-Go/config"
	"fmt"
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/user/creat", creatHandler)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil))
}
