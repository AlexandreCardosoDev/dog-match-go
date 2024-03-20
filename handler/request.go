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
	Name     string `json:"name"`
	Breed    string `json:"breed"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	OwnerID  uint   `json:"ownerid"`
}

func (r *CreatePetRequest) Validate() error {
	if r.Name == "" && r.Breed == "" && r.Birthday == "" && r.Sex == "" && r.OwnerID == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Name == "" {
		return errParamIsRequired("Name", "String")
	}
	if r.Breed == "" {
		return errParamIsRequired("Breed", "String")
	}
	if r.Birthday == "" {
		return errParamIsRequired("Birthday", "String")
	}
	if _, err := time.Parse("2006-01-02", r.Birthday); err != nil {
		return fmt.Errorf("param: %s is not a valid date, format accept YYYY-MM-DD", "Birthday")
	}
	if r.Sex == "" {
		return errParamIsRequired("Sex", "String")
	}
	if r.OwnerID == 0 {
		return errParamIsRequired("OwnerID", "Int")
	}
	return nil
}
