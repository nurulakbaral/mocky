package main

import (
	"log"
	"mocky-blog/cmd/router"
	"mocky-blog/config"
	"mocky-blog/database"
	"net/http"
)

func main() {
	db := database.New()
	router := router.New(db)
	
	log.Println("📨 Server Running on Port " + config.PORT + " .")

	http.ListenAndServe(":3001", router)
}