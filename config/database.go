package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db := "root:@tcp(127.0.0.1:3306)/toko-crud?parseTime=true"

	database, err := gorm.Open(mysql.Open(db), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung")
	}

	DB = database
}
