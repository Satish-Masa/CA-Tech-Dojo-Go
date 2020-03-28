package controllers

import (
	"CA-Tech-Dojo-Go/config"
	"fmt"
	"log"
	"net/http"
)

func Route() {
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil))
}
