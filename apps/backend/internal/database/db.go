package database

import (
	"fmt"
	"log"
	"os"
	"time" // Import time package for connection retries

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	var err error
	// Read connection details from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		host, port, user, password, dbname, sslmode, timezone)

	// Retry mechanism in case the backend starts slightly before the DB is fully ready
	maxRetries := 10
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // Log SQL queries
		})

		if err == nil {
			log.Println("Successfully connected to the database!")

			// Optional: Test the connection
			sqlDB, err := DB.DB()
			if err != nil {
				log.Printf("Failed to get underlying sql.DB: %v\n", err)
				// Decide if this is fatal or not
			} else if err = sqlDB.Ping(); err != nil {
				log.Printf("Failed to ping database: %v\n", err)
				// Decide if this is fatal or not
			}

			return // Connection successful
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		log.Printf("Retrying in %v...", retryDelay)
		time.Sleep(retryDelay)
	}

	// If loop finishes without connecting, log fatal error
	log.Fatalf("Could not connect to the database after %d attempts: %v", maxRetries, err)
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	if DB == nil {
		// This shouldn't happen if ConnectDB is called at startup,
		// but handle defensively.
		log.Println("Warning: GetDB called before ConnectDB finished or DB is nil.")
		ConnectDB() // Attempt to connect again or handle error appropriately
	}
	return DB
}
