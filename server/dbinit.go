package server

import (
	"github.com/jinzhu/gorm"
)

func migration() error {
	db, err := gormConnect()
	if err != nil {
		return err
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Portfolio{})
	return nil
}

func gormConnect() (*gorm.DB, error) {
	DBMS := "mysql"
	USER := "portfolio_app_user"
	PASS := "portfolio_password"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "portfolio_database"
	PARSETIME := "true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=" + PARSETIME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	return db, nil
}
