package domain

import "gorm.io/gorm"

type Origin struct {
	gorm.Model
	Type        OriginType `sql:"type:origin_type;not null"`
	EmployeeUID int
}
