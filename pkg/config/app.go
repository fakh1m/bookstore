package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location : ", err)
		return
	}
	now := time.Now().In(location)
	fmt.Println("Current time in Asia/Jakarta : ", now)

	dsn := "user=kimbozy password=Admin1 dbname=bookstore port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	db = d

	log.Println("Database connected!")

}

func GetDB() *gorm.DB {
	return db
}
