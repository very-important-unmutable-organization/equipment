package domain

import (
	"database/sql/driver"
)

// TODO: database needs to be prepared manually
// TODO: rubles as default

type currency string

const (
	Ruble currency = "ruble"
	USD   currency = "usd"
	Pound currency = "pound"
	Euro  currency = "euro"
)

func (p *currency) Scan(value interface{}) error {
	*p = currency(value.([]byte))
	return nil
}

func (p currency) Value() (driver.Value, error) {
	return string(p), nil
}
