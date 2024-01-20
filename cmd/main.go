package main

import (
	"context"
	"dododo/database"
	"fmt"
)

func main() {
	conn := database.NewConnect()
	defer conn.Close(context.Background())

	fmt.Println("Hello World.")
}