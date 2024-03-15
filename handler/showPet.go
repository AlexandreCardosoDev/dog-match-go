package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowPetHanler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "show pet",
	})
}
