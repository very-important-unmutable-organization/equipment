package domain

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name string
}
