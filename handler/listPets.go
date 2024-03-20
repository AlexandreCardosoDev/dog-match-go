package handler

import (
	"net/http"

	"github.com/AlexandreCardosoDev/dog-match-go/schemas"
	"github.com/gin-gonic/gin"
)

func ListPetHanler(ctx *gin.Context) {
	pet := []schemas.Pet{}

	//Get all pets
	if err := db.Find(&pet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing pets")
		return
	}

	sendSuccess(ctx, "list-pet", pet)
}
