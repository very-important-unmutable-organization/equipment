package domain

import "gorm.io/gorm"

type Purpose struct {
	gorm.Model
	Type                   PurposeType `gorm:"type:purpose_type"`
	ResponsibleEmployeeUID string      `gorm:"type:uuid;not null" json:"responsible_uid"`
}
