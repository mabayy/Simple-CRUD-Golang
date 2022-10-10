package database

import (
	"log"
	"tugas-dua/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=OrderDB port=5432 TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	db.AutoMigrate(models.Order{}, models.Item{})
	return db
}
