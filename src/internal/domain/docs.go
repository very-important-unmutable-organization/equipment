package domain

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	ItemID   uint
	Location string
}
