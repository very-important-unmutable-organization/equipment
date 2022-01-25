package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type PhotoService struct {
	photo *repository.PhotoRepo
}

func NewPhotoService(eq *repository.PhotoRepo) *PhotoService {
	return &PhotoService{photo: eq}
}

func (e *PhotoService) GetAll() (photos *[]domain.Photo, err error) {
	photos, err = e.photo.GetAll()
	return
}

func (e *PhotoService) CreatePhoto(in *domain.Photo) (photo *domain.Photo, err error) {
	if err := e.photo.CreatePhoto(in); err != nil {
		return nil, err
	}
	return in, nil
}
