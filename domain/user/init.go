package user

import "github.com/jackc/pgx/v5"

func Init(conn *pgx.Conn) Controller {
	repository := NewRepository(conn)
	service := NewService(repository)
	controller := NewController(service)
	return controller
}