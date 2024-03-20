package router

import (
	"github.com/AlexandreCardosoDev/dog-match-go/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.SetupHandler()
	basepath := "api/v1"
	v1 := router.Group(basepath)
	{
		v1.POST("/pet", handler.CreatePetHanler)
		v1.DELETE("/pet", handler.DeletePetHanler)
		v1.PUT("/pet", handler.UpdatePetHanler)
		v1.GET("/pet", handler.ShowPetHanler)
		v1.GET("pets", handler.ListPetHanler)
	}
}
