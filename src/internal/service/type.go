package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type ItemTypeService struct {
	itemType *repository.ItemTypeRepo
}

func NewItemTypeService(eq *repository.ItemTypeRepo) *ItemTypeService {
	return &ItemTypeService{itemType: eq}
}

func (e *ItemTypeService) GetAll() (itemTypes *[]domain.ItemType, err error) {
	itemTypes, err = e.itemType.GetAll()
	return
}

func (e *ItemTypeService) CreateItemType(in *domain.ItemType) (itemType *domain.ItemType, err error) {
	if err := e.itemType.CreateItemType(in); err != nil {
		return nil, err
	}
	return in, nil
}
