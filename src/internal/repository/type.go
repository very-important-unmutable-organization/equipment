package repository

import (
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type ItemTypeRepo struct {
	itemType domain.ItemType //nolint:structcheck,unused
	db       *gorm.DB
}

func NewItemTypeRepo(db *gorm.DB) *ItemTypeRepo {
	return &ItemTypeRepo{
		db: db,
	}
}

func (e *ItemTypeRepo) GetAll() (*[]domain.ItemType, error) {
	equipments := new([]domain.ItemType)
	res := e.db.Find(equipments)
	if res.Error != nil {
		return nil, res.Error
	}
	return equipments, nil
}

func (e *ItemTypeRepo) CreateItemType(itemType *domain.ItemType) (err error) {
	res := e.db.Create(itemType)
	return res.Error
}
