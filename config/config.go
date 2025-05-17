package config

import (
	"log"
	"os"
	"project-management-backend/migration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB = database
	migration.RunMigration(database)
	log.Println("Database Migration Successfully")
}
