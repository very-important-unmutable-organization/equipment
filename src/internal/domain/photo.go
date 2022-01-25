package domain

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	ItemID   uint `json:"item_id"`
	Location string
}
