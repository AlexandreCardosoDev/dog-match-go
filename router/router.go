package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	//Initialize Router
	router := gin.Default()

	//Inicialize Routes
	initializeRoutes(router)

	//Run the server
	router.Run()
}
