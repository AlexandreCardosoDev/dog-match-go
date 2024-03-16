package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name     string
	Breed    string
	Birthday time.Time
	Sex      string
	OwnerID  uint
}

type Owner struct {
	gorm.Model
	Name  string
	Email string
	Tel   string
	Pets  []Pet
}
