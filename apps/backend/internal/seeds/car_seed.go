package seeds

import (
	"log"

	"muraragi/street-racing-arena-backend/internal/database"
	"muraragi/street-racing-arena-backend/internal/models"

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

	filteredCars := database.FilterDuplicates(
		baseCars,
		models.BaseCarModel{},
		db.Model(&models.BaseCarModel{}),
		func(item any) string {
			if car, ok := item.(models.BaseCarModel); ok {
				return car.Name
			}
			return ""
		},
	)

	if len(filteredCars) > 0 {
		log.Printf("Unique base cars to insert: %+v\n", filteredCars)
		if err := db.Model(&models.BaseCarModel{}).Create(&filteredCars).Error; err != nil {
			log.Fatalf("Failed to batch seed base cars: %v", err)
		}

		log.Println("BaseCarModel seeding complete.")
	}
}
