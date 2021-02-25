package server

import (
	"fmt"

	"github.com/jinzhu/gorm"
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
	DBMS := "postgres"
	HOST := "localhost"
	PORT := "5432"
	USER := "app_user"
	PASS := "teshimapw"
	DBNAME := "prototype_database"

	CONNECT := "host=" + HOST + " port=" + PORT + " user=" + USER + " dbname=" + DBNAME + " password=" + PASS + " sslmode=disable"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.SingularTable(true)
	return db, nil
}
