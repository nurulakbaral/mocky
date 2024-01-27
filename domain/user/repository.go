package user

import (
	"context"
	"dododo/logger"

	"gorm.io/gorm"
)


type Repository interface {
	Insert(ctx context.Context, user User) (User, error)
	SelectById(ctx context.Context, id string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{ db }
}

func (repo *repository) Insert(ctx context.Context, user User) (User, error) {
	results := repo.db.Create(&user)
	if results.Error != nil {
		logger.Logger("Can not create Insert user.", results.Error)
	}
	return user, results.Error
}	

func (repo *repository) SelectById(ctx context.Context, id string) (User, error) {
	return User{}, nil
}