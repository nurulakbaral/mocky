package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)


type Repository interface {
	Insert(ctx context.Context, user User)
	SelectById(ctx context.Context, id string) (User, error)
}

type repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) Repository {
	return &repository{ conn }
}

func (repo *repository) Insert(ctx context.Context, user User) {}

func (repo *repository) SelectById(ctx context.Context, id string) (User, error) {
	return User{}, nil
}