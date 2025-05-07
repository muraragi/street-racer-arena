package main

import (
	"log"
	"net/http"

	"muraragi/street-racer-arena-backend/internal/database"
	"muraragi/street-racer-arena-backend/internal/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: Could not load .env file:", err)
	}

	database.ConnectDB()
	database.MigrateDB()

	db := database.GetDB()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello, World!!"})
	})

	router.GET("/cars", func(c *gin.Context) {
		var cars []models.BaseCarModel
		db.Find(&cars)
		c.IndentedJSON(http.StatusOK, cars)
	})

	router.Run(":8080")
}
