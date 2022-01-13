package repository

import (
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type Equipment struct {
	db *gorm.DB
}

func NewEquipment(db *gorm.DB) Equipment {
	return Equipment{
		db: db,
	}
}

func (e *Equipment) GetAll() (*[]domain.Equipment, error) {
	equipments := new([]domain.Equipment)
	res := e.db.Find(equipments)
	if res.Error != nil {
		return nil, res.Error
	}
	return equipments, nil
}

func (e *Equipment) Post(equipment *domain.Equipment) error {
	res := e.db.Create(equipment)
	return res.Error
}
