package service

import (
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
)

type StateService struct {
	state *repository.StateRepo
}

func NewStateService(eq *repository.StateRepo) *StateService {
	return &StateService{state: eq}
}

func (e *StateService) GetAll() (states *[]domain.State, err error) {
	states, err = e.state.GetAll()
	return
}

func (e *StateService) CreateState(in *domain.State) (state *domain.State, err error) {
	if err := e.state.CreateState(in); err != nil {
		return nil, err
	}
	return in, nil
}
