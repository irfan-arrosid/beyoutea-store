package migration

import (
	"fmt"
	"kstyle-be-techtest/database"
	"kstyle-be-techtest/models/entities"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&entities.Member{},
		&entities.Product{},
		&entities.Review_product{},
	)

	if err != nil {
		log.Println("Database migration is failed")
	} else {
		fmt.Println("Database migrated")
	}
}
