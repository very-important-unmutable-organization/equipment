package domain

import "gorm.io/gorm"

type ItemType struct {
	gorm.Model
	Category string `sql:"type:varchar;not null"`
	Name     string `sql:"type:varchar;not null"`
}
