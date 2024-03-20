package handler

import (
	"github.com/AlexandreCardosoDev/dog-match-go/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func SetupHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSql()
}
