package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)


func NewConnect() *pgx.Conn  {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Unable to connect to database: \n", err)
	}
	return conn
}