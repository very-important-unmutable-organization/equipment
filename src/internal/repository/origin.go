package repository

import (
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type OriginRepo struct {
	origin domain.Origin //nolint:structcheck,unused
	db     *gorm.DB
}

func NewOriginRepo(db *gorm.DB) *OriginRepo {
	return &OriginRepo{
		db: db,
	}
}

func (e *OriginRepo) GetAll() (*[]domain.Origin, error) {
	origins := new([]domain.Origin)
	res := e.db.Find(origins)
	if res.Error != nil {
		return nil, res.Error
	}
	return origins, nil
}

func (e *OriginRepo) CreateOrigin(origin *domain.Origin) (err error) {
	res := e.db.Create(origin)
	return res.Error
}
