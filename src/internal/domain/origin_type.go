package domain

import (
	"database/sql/driver"
)

type OriginType string

const (
	CompanyProperty  OriginType = "company"
	EmployeeProperty OriginType = "employee"
)

func (p *OriginType) Scan(value interface{}) error {
	*p = OriginType(value.(string))
	return nil
}

func (p OriginType) Value() (driver.Value, error) {
	return string(p), nil
}
