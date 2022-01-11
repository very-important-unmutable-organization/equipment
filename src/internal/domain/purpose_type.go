package domain

import (
	"database/sql/driver"
)

// TODO: database needs to be prepared manually

type PurposeType string

const (
	Personal PurposeType = "personal"
	General  PurposeType = "general"
	Testing  PurposeType = "testing"
)

func (p *PurposeType) Scan(value interface{}) error {
	*p = PurposeType(value.([]byte))
	return nil
}

func (p PurposeType) Value() (driver.Value, error) {
	return string(p), nil
}
