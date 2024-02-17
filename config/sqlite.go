package config

import (
	"os"

	"github.com/vapdev/gopportunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if DB exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// Create DB
		logger.Info("Creating SQLite DB")
		// Create directory if it doesn't exist
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating SQLite: %v", err)
			return nil, err
		}
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to SQLite: %v", err)
		return nil, err
	}
	// Migrate schema
	err = db.AutoMigrate(&schemas.Opportunity{})
	if err != nil {
		logger.Errorf("Error migrating schema: %v", err)
		return nil, err
	}
	return db, nil
}
