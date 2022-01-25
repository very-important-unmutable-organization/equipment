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

func (e *EquipmentRepo) EditById(id int, equipment *domain.Equipment) (err error) {
	existing := new(domain.Equipment)
	res := e.db.First(existing, id)
	if res.Error != nil {
		return res.Error
	}
	equipment.ID = existing.ID
	res = e.db.Save(equipment)
	return res.Error
}

func (e *EquipmentRepo) Take(id int) error {
	equipment := new(domain.Equipment)
	res := e.db.First(equipment, id)
	if res.Error != nil {
		return res.Error
	}
	if equipment.Status == domain.Taken {
		return domain.ErrorEquipmentTaken{}
	}
	equipment.Status = domain.Taken
	e.db.Save(equipment)
	return nil
}

func (e *EquipmentRepo) Free(id int) error {
	equipment := new(domain.Equipment)
	res := e.db.First(equipment, id)
	if res.Error != nil {
		return res.Error
	}
	if equipment.Status == domain.Free {
		return domain.ErrorEquipmentFree{}
	}
	equipment.Status = domain.Free
	e.db.Save(equipment)
	return nil
}
