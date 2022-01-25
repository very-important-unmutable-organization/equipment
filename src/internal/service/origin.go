package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type OriginService struct {
	origin *repository.OriginRepo
}

func NewOriginService(eq *repository.OriginRepo) *OriginService {
	return &OriginService{origin: eq}
}

func (e *OriginService) GetAll() (origins *[]domain.Origin, err error) {
	origins, err = e.origin.GetAll()
	return
}

func (e *OriginService) CreateOrigin(in *domain.Origin) (origin *domain.Origin, err error) {
	if err := e.origin.CreateOrigin(in); err != nil {
		return nil, err
	}
	return in, nil
}
