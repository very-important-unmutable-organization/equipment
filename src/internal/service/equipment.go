package service

import (
	"strconv"

	qrcode "github.com/skip2/go-qrcode"

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

func (e *EquipmentService) EditById(id int, equipment *domain.Equipment) error {
	err := e.equipment.EditById(id, equipment)
	return err
}

func (e *EquipmentService) Take(id int) error {
	return e.equipment.Take(id)
}

func (e *EquipmentService) Free(id int) error {
	return e.equipment.Free(id)
}

func (e *EquipmentService) GetQrForId(id int, prefix string) ([]byte, error) {
	_, err := e.equipment.GetById(id)
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(prefix+strconv.Itoa(id), qrcode.Medium, 256)
}
