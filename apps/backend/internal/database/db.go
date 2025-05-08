package database

import (
	"fmt"
	"log"
	"muraragi/street-racer-arena-backend/internal/models"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		host, port, user, password, dbname, sslmode, timezone)

	maxRetries := 10
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err == nil {
			log.Println("Successfully connected to the database!")

			sqlDB, err := DB.DB()
			if err != nil {
				log.Printf("Failed to get underlying sql.DB: %v\n", err)
			} else if err = sqlDB.Ping(); err != nil {
				log.Printf("Failed to ping database: %v\n", err)
			}

			return
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		log.Printf("Retrying in %v...", retryDelay)
		time.Sleep(retryDelay)
	}

	log.Fatalf("Could not connect to the database after %d attempts: %v", maxRetries, err)
}

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.BaseCarModel{},
		&models.Car{},
		&models.CarComponent{},
		&models.InstalledCarComponent{},
		&models.Track{},
		&models.LobbyTrack{},
		&models.Lobby{},
		&models.LobbyParticipant{},
		&models.LobbyResult{},
		&models.LobbyResultEntry{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Successfully migrated database.")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Println("Warning: GetDB called before ConnectDB finished or DB is nil.")
		ConnectDB()
	}
	return DB
}

func FilterDuplicates[T any, K comparable, M any](
	items []T,
	dbModel M,
	query *gorm.DB,
	extractKey func(item interface{}) K,
) []T {
	if len(items) == 0 {
		return items
	}

	var existingItems []M
	if err := query.Find(&existingItems).Error; err != nil {
		log.Printf("Error querying database in FilterDuplicates: %v", err)
		return items
	}

	existingKeys := make(map[K]struct{})
	for _, item := range existingItems {
		key := extractKey(item)
		existingKeys[key] = struct{}{}
	}

	var filteredItems []T
	for _, item := range items {
		key := extractKey(item)
		if _, exists := existingKeys[key]; !exists {
			filteredItems = append(filteredItems, item)
		}
	}

	return filteredItems
}
