package repository

import (
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type EquipmentRepo struct {
	equipment domain.Equipment //nolint:structcheck,unused
	db        *gorm.DB
}

func NewEquipmentRepo(db *gorm.DB) *EquipmentRepo {
	return &EquipmentRepo{
		db: db,
	}
}

func (e *EquipmentRepo) GetAll() (*[]domain.Equipment, error) {
	equipments := new([]domain.Equipment)
	res := e.db.Find(equipments)
	if res.Error != nil {
		return nil, res.Error
	}
	return equipments, nil
}

func (e *EquipmentRepo) CreateEquipment(equipment *domain.Equipment) (err error) {
	res := e.db.Create(equipment)
	return res.Error
}

func (e *EquipmentRepo) GetById(id int) (*domain.Equipment, error) {
	equipment := new(domain.Equipment)
	res := e.db.First(equipment, id)
	return equipment, res.Error
}
