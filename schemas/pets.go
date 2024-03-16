package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Name  string
	Email string
	Tel   string
	Pets  []Pet
}

type Pet struct {
	gorm.Model
	Name      string
	Breed     string
	Birthday  time.Time
	Sex       string
	Available bool
	OwnerID   uint
}
