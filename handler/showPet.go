package handler

import (
	"fmt"
	"net/http"

	"github.com/AlexandreCardosoDev/dog-match-go/schemas"
	"github.com/gin-gonic/gin"
)

func ShowPetHanler(ctx *gin.Context) {
	id := ctx.Query("id")

	//Validate if id is valid
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	pet := schemas.Pet{}
	//Find Pet ID
	if err := db.First(&pet, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("pet with id: %v not found", id))
		return
	}
	//return pet
	sendSuccess(ctx, "show-pet", pet)
}
