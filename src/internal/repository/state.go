package repository

import (
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

type StateRepo struct {
	state domain.State //nolint:structcheck,unused
	db    *gorm.DB
}

func NewStateRepo(db *gorm.DB) *StateRepo {
	return &StateRepo{
		db: db,
	}
}

func (e *StateRepo) GetAll() (*[]domain.State, error) {
	states := new([]domain.State)
	res := e.db.Find(states)
	if res.Error != nil {
		return nil, res.Error
	}
	return states, nil
}

func (e *StateRepo) CreateState(state *domain.State) (err error) {
	res := e.db.Create(state)
	return res.Error
}
