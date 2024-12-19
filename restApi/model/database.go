package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Database() (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gormpass dbname=gorm port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&Grocery{}); err != nil {
		log.Fatal(err)
	}
	return db, err
}
