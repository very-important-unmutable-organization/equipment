package domain

import "gorm.io/gorm"

type Origin struct {
	gorm.Model
	Type        OriginType `gorm:"type:origin_type;not null"`
	EmployeeUID int        `json:"employee_uid"`
}
