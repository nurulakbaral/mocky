package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		log.Fatalf("Error executing query: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatalf("Error scanning row: %v\n", err)
		}

		fmt.Printf("ID: %d, Name: %s, OtherColumn: %v\n", id, name, age)
	}

}