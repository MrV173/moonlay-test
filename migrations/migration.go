package migrations

import (
	"fmt"
	"moonlay-test/models"
	"moonlay-test/pkg/database"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&models.List{},
		&models.Sublist{})

	if err != nil {
		fmt.Println(err)
		panic(("Migration Failed"))
	}

	fmt.Println(("Migration Success"))
}
