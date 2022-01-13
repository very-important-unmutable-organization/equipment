package domain

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	ItemID   uint `json:"item_id"`
	Location string
}
