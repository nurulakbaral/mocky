package user

import (
	"time"

	"gorm.io/gorm"
)


type User struct {
	Id string  `gorm:"primaryKey"`
	Username string  `gorm:"unique"`
	Email string `gorm:"unique"`
	Password string `gorm:"not null"`
	CreatedAt time.Time
  	UpdatedAt time.Time
  	DeletedAt gorm.DeletedAt
}