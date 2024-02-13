package main

import (
	"cc_terminal/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=rishi password=1111 dbname=cc_terminal port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	return db, nil
}

func main() {
	db, err := InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the User model
	db.AutoMigrate(&models.User{})
}
