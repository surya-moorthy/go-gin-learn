package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("./database.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&Grocery{}); err != nil {
		log.Fatal(err)
	}
	return db, err
}
