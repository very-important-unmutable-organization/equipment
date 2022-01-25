package repository

import (
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type PurposeRepo struct {
	purpose domain.Purpose //nolint:structcheck,unused
	db      *gorm.DB
}

func NewPurposeRepo(db *gorm.DB) *PurposeRepo {
	return &PurposeRepo{
		db: db,
	}
}

func (e *PurposeRepo) GetAll() (*[]domain.Purpose, error) {
	purposes := new([]domain.Purpose)
	res := e.db.Find(purposes)
	if res.Error != nil {
		return nil, res.Error
	}
	return purposes, nil
}

func (e *PurposeRepo) CreatePurpose(purpose *domain.Purpose) (err error) {
	res := e.db.Create(purpose)
	return res.Error
}
