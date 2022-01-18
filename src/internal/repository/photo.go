package repository

import (
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type PhotoRepo struct {
	photo domain.Photo //nolint:structcheck,unused
	db    *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) *PhotoRepo {
	return &PhotoRepo{
		db: db,
	}
}

func (e *PhotoRepo) GetAll() (*[]domain.Photo, error) {
	photos := new([]domain.Photo)
	res := e.db.Find(photos)
	if res.Error != nil {
		return nil, res.Error
	}
	return photos, nil
}

func (e *PhotoRepo) CreatePhoto(photo *domain.Photo) (err error) {
	res := e.db.Create(photo)
	return res.Error
}
