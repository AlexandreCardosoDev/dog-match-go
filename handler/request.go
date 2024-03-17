package handler

import (
	"fmt"
	"time"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

//CreatePets

type CreatePetRequest struct {
	Name     string    `json:"name"`
	Breed    string    `json:"breed"`
	Birthday time.Time `json:"brithday"`
	Sex      string    `json:"sex"`
	OwnerID  uint      `json:"ownerid"`
}

func (r *CreatePetRequest) Validate() error {
	if r.Name == "" && r.Breed == "" && r.Sex == "" && r.OwnerID == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Name == "" {
		return errParamIsRequired("Name", "String")
	}
	if r.Breed == "" {
		return errParamIsRequired("Breed", "String")
	}
	if r.Sex == "" {
		return errParamIsRequired("Sex", "String")
	}
	return nil
}
