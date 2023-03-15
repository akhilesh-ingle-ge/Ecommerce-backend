package config

import (
	"fmt"
	"os"

	"github.com/ecom/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetUpDb() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("Failed to load Database")
	}
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{})
	DB = db
	fmt.Println("Connected to database")
}
