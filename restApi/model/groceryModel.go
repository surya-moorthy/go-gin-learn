package model

import (
	"gorm.io/gorm"
)

type Grocery struct {

	gorm.Model
	Name string `json:"name"`
	Quantity string `json:"Quantity"`
}

