package main

import (
	"context"
	"dododo/cmd/router"
	"dododo/config"
	"dododo/database"
	"log"
	"net/http"
)

func main() {
	conn := database.NewConnect()
	defer conn.Close(context.Background())

	router := router.New(conn)
	log.Println("📨 Server Running on Port " + config.Port + " .")

	http.ListenAndServe(":3001", router)
}