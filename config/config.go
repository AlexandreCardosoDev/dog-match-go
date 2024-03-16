package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitConfig() error {
	var err error
	// Initialize SQLite
	db, err = InitializeSqlLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite %v", err)
	}
	return nil
}

func GetSql() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
