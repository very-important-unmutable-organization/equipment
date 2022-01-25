package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type Services struct {
	EquipmentService *EquipmentService
	ItemTypeService  *ItemTypeService
	StateService     *StateService
	PurposeService   *PurposeService
	OriginService    *OriginService
	DocumentService  *DocumentService
	PhotoService     *PhotoService
}

func NewServices(repos *repository.Repositories) (*Services, error) {
	return &Services{
		EquipmentService: NewEquipmentService(repos.EquipmentRepo),
		ItemTypeService:  NewItemTypeService(repos.ItemTypeRepo),
		StateService:     NewStateService(repos.StateRepo),
		PurposeService:   NewPurposeService(repos.PurposeRepo),
		OriginService:    NewOriginService(repos.OriginRepo),
		DocumentService:  NewDocumentService(repos.DocumentRepo),
		PhotoService:     NewPhotoService(repos.PhotoRepo),
	}, nil
}
