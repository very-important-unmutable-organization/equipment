package repository

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Database string
	Password string
}

type Repositories struct {
	Equipment domain.Equipment
}

func NewRepositories(cfg Config) (*Repositories, error) {
	return &Repositories{}, nil
}
