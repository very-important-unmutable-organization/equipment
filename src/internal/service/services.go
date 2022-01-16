package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type Services struct {
	EquipmentService *EquipmentService
}

func NewServices(repos *repository.Repositories) (*Services, error) {
	equipmentService := NewEquipmentService(*repos.EquipmentRepo)

	return &Services{EquipmentService: equipmentService}, nil
}
