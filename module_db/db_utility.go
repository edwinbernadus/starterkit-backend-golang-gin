package module_db

import (
	"starterkit/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnString() string {
	dsn := "host=localhost user=edwin password=Testing1 dbname=golang_starterkit port=5432 sslmode=disable"
	return dsn
}

func GetDb() *gorm.DB {

	var dsn = GetConnString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	var _ = db.AutoMigrate(&models.Student{})

	return db
}
