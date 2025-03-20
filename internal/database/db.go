package database

import (
	"log"
	"os"

	"github.com/JonathanGeorgiou/astrolog/internal/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbPath := os.Getenv("HOME") + "/.astrolog/astrolog.db"
	os.MkdirAll(os.Getenv("HOME")+"/.astrolog", os.ModePerm)
	var err error

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the Task model
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database initialized successfully!")
}
