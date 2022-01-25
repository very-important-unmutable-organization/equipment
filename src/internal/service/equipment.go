package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type EquipmentService struct {
	equipment *repository.EquipmentRepo
}

func NewEquipmentService(eq *repository.EquipmentRepo) *EquipmentService {
	return &EquipmentService{equipment: eq}
}

func (e *EquipmentService) GetAll() (equipments *[]domain.Equipment, err error) {
	equipments, err = e.equipment.GetAll()
	return
}

func (e *EquipmentService) CreateEquipment(in *domain.Equipment) (equipment *domain.Equipment, err error) {
	if err := e.equipment.CreateEquipment(in); err != nil {
		return nil, err
	}
	return in, nil
}

func (e *EquipmentService) GetById(id int) (*domain.Equipment, error) {
	equipment, err := e.equipment.GetById(id)
	return equipment, err
}
