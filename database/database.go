package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
	env := godotenv.Load()
	if env != nil {
		panic("Failed to load env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	var err error
	connect := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		dbUser, dbPass, dbHost, dbPort, dbName)

	DB, err = gorm.Open(sqlserver.Open(connect), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

}
