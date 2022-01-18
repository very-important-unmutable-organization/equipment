package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type PurposeService struct {
	purpose *repository.PurposeRepo
}

func NewPurposeService(eq *repository.PurposeRepo) *PurposeService {
	return &PurposeService{purpose: eq}
}

func (e *PurposeService) GetAll() (purposes *[]domain.Purpose, err error) {
	purposes, err = e.purpose.GetAll()
	return
}

func (e *PurposeService) CreatePurpose(in *domain.Purpose) (purpose *domain.Purpose, err error) {
	if err := e.purpose.CreatePurpose(in); err != nil {
		return nil, err
	}
	return in, nil
}
