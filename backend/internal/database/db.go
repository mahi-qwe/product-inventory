package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"product-inventory/internal/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=9675 dbname=inventory port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	db.AutoMigrate(
		&models.Product{},
		&models.Variant{},
		&models.VariantOption{},
		&models.SubVariant{},
		&models.StockTransaction{},
	)
}
