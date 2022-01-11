package domain

import "gorm.io/gorm"

type Purpose struct {
	gorm.Model
	Type                   PurposeType `sql:"type:purpose_type"`
	ResponsibleEmployeeUID uint        `sql:"not null"`
}
