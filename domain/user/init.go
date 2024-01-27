package user

import (
	"gorm.io/gorm"
)

func Init(db *gorm.DB) Controller {
	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)
	return controller
}