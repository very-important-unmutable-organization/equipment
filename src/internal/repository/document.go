package repository

import (
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type DocumentRepo struct {
	document domain.Document //nolint:structcheck,unused
	db       *gorm.DB
}

func NewDocumentRepo(db *gorm.DB) *DocumentRepo {
	return &DocumentRepo{
		db: db,
	}
}

func (e *DocumentRepo) GetAll() (*[]domain.Document, error) {
	documents := new([]domain.Document)
	res := e.db.Find(documents)
	if res.Error != nil {
		return nil, res.Error
	}
	return documents, nil
}

func (e *DocumentRepo) CreateDocument(document *domain.Document) (err error) {
	res := e.db.Create(document)
	return res.Error
}
