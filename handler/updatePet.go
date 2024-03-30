package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AlexandreCardosoDev/dog-match-go/schemas"
	"github.com/gin-gonic/gin"
)

func UpdatePetHanler(ctx *gin.Context) {
	id := ctx.Query("id")
	//Validate if id isn't empty
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	pet := schemas.Pet{}
	//Validate fields from request
	request := CreatePetRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//Find pet in database
	if err := db.First(&pet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("pet with id: %v not found", id))
		return
	}
	birthday, _ := time.Parse("2006-01-02", request.Birthday)
	pet.Name = request.Name
	pet.Breed = request.Breed
	pet.Sex = request.Sex
	pet.Birthday = birthday
	pet.OwnerID = request.OwnerID

	if err := db.Save(&pet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(ctx, "update-Pet", pet)
}
