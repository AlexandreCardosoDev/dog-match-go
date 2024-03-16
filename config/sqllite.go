package config

import (
	"os"

	"github.com/AlexandreCardosoDev/dog-match-go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/pets.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating ...")
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initialize db %v", err)
		return nil, err
	}
	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Pet{}, &schemas.Owner{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
