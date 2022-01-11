package domain

import "gorm.io/gorm"

type State struct {
	gorm.Model
	Name string
}
