package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnection() {
	dsn := "irfanarrosid:my04sql04@tcp(localhost:4306)/kstyle_be?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection database is failed.")
	} else {
		fmt.Println("Database connected....")
	}

	DB = db
}
