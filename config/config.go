package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetDB() (*gorm.DB, error) {
	var err error
	dsn := "host=localhost user=postgres password=Berandallasta15 dbname=db_warehouse port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
		log.Fatal("Failed connect to database")
	}
	log.Println("Success connect to database")

	return db, nil
}
