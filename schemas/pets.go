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

type PetResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Breed     string    `json:"breed"`
	Birthday  time.Time `json:"birthday"`
	Sex       string    `json:"sex"`
	OwnerID   uint      `json:"ownerId"`
}

type OwnerResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Tel       string    `json:"tel"`
}
