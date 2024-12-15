package config

import (
	"log"
	"project-management-backend/migration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user=archit password=archit dbname=project_management port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB = database
	migration.RunMigration(database)
	log.Println("Database Migration Successfully")
}
