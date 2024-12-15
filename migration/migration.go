package migration

import (
	"log"
	"project-management-backend/models"

	"gorm.io/gorm"
)

func RunMigration(database *gorm.DB) {
	modelsMigrate := []interface{}{
		&models.User{},
		&models.LeaveRecord{},
		&models.LeaveType{},
	}
	for _, model := range modelsMigrate {
		err := database.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Failed to migrate %T:%v", model, err)
		}
		log.Printf("%T migrated Successfully", model)
	}
}
