package config

import (
	"assignment/app/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbUser := "postgres"
	dbPass := "123456"
	dbHost := "localhost"
	dbName := "assignment_two"
	dbPort := "5432"

	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect postgres database")
	}

	return db
}

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(
		&models.Order{}, &models.Item{})
	fmt.Println("Database Migration Completed!")
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
