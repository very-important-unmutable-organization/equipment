package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type DocumentService struct {
	document *repository.DocumentRepo
}

func NewDocumentService(eq *repository.DocumentRepo) *DocumentService {
	return &DocumentService{document: eq}
}

func (e *DocumentService) GetAll() (documents *[]domain.Document, err error) {
	documents, err = e.document.GetAll()
	return
}

func (e *DocumentService) CreateDocument(in *domain.Document) (document *domain.Document, err error) {
	if err := e.document.CreateDocument(in); err != nil {
		return nil, err
	}
	return in, nil
}
