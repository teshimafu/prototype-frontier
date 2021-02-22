package server

import (
	"github.com/jinzhu/gorm"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "portfolio_app_user"
	PASS := "teshima"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "portfolio_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	db.SingularTable(true)
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Portfolio{})
	return db
}
