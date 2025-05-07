package seeds

import (
	"log"

	"muraragi/street-racer-arena-backend/internal/models"

	"gorm.io/gorm"
)

func SeedCars(db *gorm.DB) {
	baseCars := []models.BaseCarModel{
		{Name: "Toyota Supra", BasePower: 220, BaseHandling: 130},
		{Name: "Nissan GT-R", BasePower: 250, BaseHandling: 140},
		{Name: "Mazda RX-7", BasePower: 200, BaseHandling: 150},
		{Name: "Ford Mustang GT", BasePower: 260, BaseHandling: 110},
		{Name: "Chevrolet Camaro SS", BasePower: 245, BaseHandling: 115},
		{Name: "Honda Civic Type R", BasePower: 180, BaseHandling: 160},
		{Name: "Subaru Impreza WRX STI", BasePower: 210, BaseHandling: 155},
		{Name: "BMW M3", BasePower: 230, BaseHandling: 145},
	}

	for _, bc := range baseCars {
		var existing models.BaseCarModel
		if err := db.First(&existing, "name = ?", bc.Name).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&bc).Error; err != nil {
				log.Fatalf("Failed to seed base car %s: %v", bc.Name, err)
			}
			log.Printf("Seeded BaseCarModel: %s", bc.Name)
		} else if err != nil {
			log.Fatalf("Error checking for existing base car %s: %v", bc.Name, err)
		}
	}
	log.Println("BaseCarModel seeding complete.")
}
