package handler

import (
	"net/http"
	"time"

	"github.com/AlexandreCardosoDev/dog-match-go/schemas"
	"github.com/gin-gonic/gin"
)

func CreatePetHanler(ctx *gin.Context) {
	request := CreatePetRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	birthdayRequest, _ := time.Parse("2006-01-02", request.Birthday)

	pet := schemas.Pet{
		Name:     request.Name,
		Breed:    request.Breed,
		Birthday: birthdayRequest,
		Sex:      request.Sex,
		OwnerID:  request.OwnerID,
	}

	if err := db.Create(&pet).Error; err != nil {
		logger.Errorf("error creating pet %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating pet on database")
		return
	}

	sendSuccess(ctx, "create-pet", pet)
}
