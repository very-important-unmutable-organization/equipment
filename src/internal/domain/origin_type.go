package domain

import (
	"database/sql/driver"
)

// TODO: database needs to be prepared manually

type OriginType string

const (
	CompanyProperty  OriginType = "company"
	EmployeeProperty OriginType = "employee"
)

func (p *OriginType) Scan(value interface{}) error {
	*p = OriginType(value.([]byte))
	return nil
}

func (p OriginType) Value() (driver.Value, error) {
	return string(p), nil
}
