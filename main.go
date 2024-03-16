package main

import (
	"github.com/AlexandreCardosoDev/dog-match-go/config"
	"github.com/AlexandreCardosoDev/dog-match-go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.InitConfig()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.SetupRouter()
}
