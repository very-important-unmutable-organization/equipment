package domain

import (
	"database/sql/driver"
)

type Status string

const (
	Free  Status = "free"
	Taken Status = "taken"
)

func (p *Status) Scan(value interface{}) error {
	*p = Status(value.(string))
	return nil
}

func (p Status) Value() (driver.Value, error) {
	return string(p), nil
}
