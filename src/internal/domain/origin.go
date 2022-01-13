package domain

import "gorm.io/gorm"

type Origin struct {
	gorm.Model
	Type        OriginType `gorm:"type:origin_type;not null"`
	EmployeeUID string     `gorm:"type:uuid" json:"employee_uid"`
}
