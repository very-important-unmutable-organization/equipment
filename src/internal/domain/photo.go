package domain

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	ItemID   uint
	Location string
}
