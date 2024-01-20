package user

import (
	"context"
)

type Service interface{
	Register(ctx context.Context, user User) (User, error)
	Login(ctx context.Context, user User) (User, error)
	Logout(ctx context.Context)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (service *service) Register(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

func (service *service) Login(ctx context.Context, user User) (User, error) {
		return User{}, nil
}

func  (service *service) Logout(ctx context.Context) {}
