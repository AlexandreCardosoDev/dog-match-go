package handler

import (
	"fmt"
	"net/http"

	"github.com/AlexandreCardosoDev/dog-match-go/schemas"
	"github.com/gin-gonic/gin"
)

func DeletePetHanler(ctx *gin.Context) {
	id := ctx.Query("id")

	//Validate if id has a number greater than 0
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	pet := schemas.Pet{}
	//Find Pet
	if err := db.First(&pet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("pet with id: %s not found", id))
		return
	}

	//Delete Pet
	if err := db.Delete(&pet).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting pet with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-pet", pet)

}
