package migrations

import (
	"crud-golang/crud/config"
	"crud-golang/crud/src/database"
	"log"
)

func RunMigration() {
	db := config.NewDB()
	// auto migrtae with gorm for new table database create schema in folder database in file database-schema.go
	err := db.AutoMigrate(&database.Users{}, &database.CategoryBook{}, &database.Book{})
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
	log.Println("migrate completed successfully")
}
