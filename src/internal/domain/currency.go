package domain

import (
	"database/sql/driver"
)

type Currency string

const (
	Ruble Currency = "ruble"
	USD   Currency = "usd"
	Pound Currency = "pound"
	Euro  Currency = "euro"
)

func (p *Currency) Scan(value interface{}) error {
	*p = Currency(value.(string))
	return nil
}

func (p Currency) Value() (driver.Value, error) {
	return string(p), nil
}
