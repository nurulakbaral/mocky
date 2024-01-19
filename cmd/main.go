package main

import (
	"context"
	"dododo/db"
	"fmt"
)

func main() {
	db := db.New()
	defer db.Close(context.Background())

	fmt.Println("Hello World.")
}