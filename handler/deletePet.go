package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePetHanler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "delete pet",
	})
}
