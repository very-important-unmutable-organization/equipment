package repository

import (
	"fmt"
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func InitDb(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(fmt.Sprintf("could not connect to database: %s", err))
	}

	if err = db.AutoMigrate(&domain.Equipment{}); err != nil {
		panic(err)
	}

	return db, nil
}

func NewRepositories(cfg Config) (*Repositories, error) {
	return &Repositories{}, nil
}
