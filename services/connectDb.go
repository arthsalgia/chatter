package services

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDb(file string) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 2 * time.Second,
			LogLevel:      logger.Warn,
		},
	)
	if file == "" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			log.Fatalf("Failed to detect home directory: %v", err)
		}

		dbPath := filepath.Join(homeDir, "Library", "Messages", "chat.db")

		DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: newLogger})

		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		fmt.Println("Successfully connected to database at ~/Library/Messages/chat.db")
	} else {
		var err error
		DB, err = gorm.Open(sqlite.Open(file), &gorm.Config{Logger: newLogger})
		if err != nil {

			log.Fatalf("Failed to connect to database: %v", err)
		}

		fmt.Println("Successfully connected to database at", file)
	}
}
