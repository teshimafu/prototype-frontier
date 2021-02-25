package server

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func migration() error {
	db, err := gormConnect()
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&Portfolio{})
	return nil
}

func gormConnect() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		HOST := os.Getenv("POSTGRES_HOST")
		PORT := os.Getenv("POSTGRES_PORT")
		USER := os.Getenv("POSTGRES_USER")
		PASS := os.Getenv("POSTGRES_PASSWORD")
		DBNAME := os.Getenv("POSTGRES_DBNAME")
		databaseURL = "host=" + HOST + " port=" + PORT + " user=" + USER + " dbname=" + DBNAME + " password=" + PASS + " sslmode=disable"
	}
	db, err := gorm.Open("postgres", databaseURL)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.SingularTable(true)
	return db, nil
}
