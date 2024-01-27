package main

import (
	"dododo/cmd/router"
	"dododo/config"
	"dododo/database"
	"log"
	"net/http"
)

func main() {
	db := database.New()
	router := router.New(db)
	
	log.Println("📨 Server Running on Port " + config.PORT + " .")

	http.ListenAndServe(":3001", router)
}