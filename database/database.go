package database

import (
	"fmt"
	"strconv"

	"echo.osaxon/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.Atoi(p)

	if err != nil {
		panic("Invalid port number")
	}

	dsn := fmt.Sprintf(
		"host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)

	println(dsn)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Print(DB)
	fmt.Println("Connection Opened to Database")

	return DB, nil
}
