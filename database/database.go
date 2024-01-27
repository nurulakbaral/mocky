package database

import (
	"dododo/config"
	"dododo/domain/user"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewMigration() {
	err := db.AutoMigrate(&user.User{})
	if err != nil {
		log.Println("Failed to migrate tables.")
		panic(err)
	}

}

func New() *gorm.DB  {
	var err error
	db, err = gorm.Open(postgres.Open(config.DATABASE_URL), &gorm.Config{})

	if err != nil {
		log.Println("Unable to connect to database: \n", err)
	}

	NewMigration()

	log.Println("⚙️  Connected to The Database.")
	
	return db
}