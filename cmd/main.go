package main

import (
	"context"
	"dododo/cmd/router"
	"dododo/database"
	"log"
	"net/http"
)

func main() {
	conn := database.NewConnect()
	defer conn.Close(context.Background())

	router := router.New(conn)
	log.Println("📨 Server Running on Port 3001")

	http.ListenAndServe(":3001", router)
}