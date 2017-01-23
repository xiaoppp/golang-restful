package main

import (
	"golang-restful/boot"
	"golang-restful/routers"
	"log"
	"net/http"
)

func main() {
	boot.Start()

	router := routers.Config()

	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}
