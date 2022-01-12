package domain

import "gorm.io/gorm"

type ItemType struct {
	gorm.Model
	Category string `gorm:"type:varchar;not null"`
	Name     string `gorm:"type:varchar;not null"`
}
