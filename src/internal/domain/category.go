package domain

import (
	"database/sql/driver"
)

// TODO: database needs to be prepared manually

type Category string

const (
	Furniture         Category = "furniture"
	OfficeEquipment   Category = "office_equipment"
	PersonalEquipment Category = "personal_equipment"
)

func (p *Category) Scan(value interface{}) error {
	*p = Category(value.([]byte))
	return nil
}

func (p Category) Value() (driver.Value, error) {
	return string(p), nil
}
