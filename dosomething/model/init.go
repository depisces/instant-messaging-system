package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql数据库连接失败", err)
	}
	DB = db
	migration()
}
