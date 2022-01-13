package rest

import (
	"github.com/go-chi/chi"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
	"github.com/very-important-unmutable-organization/equipment/internal/service"
	"gorm.io/gorm"
)

type Equipment struct {
}

func (e Equipment) RegisterEndpoints(r *chi.Mux, db *gorm.DB) {
	rep := repository.NewEquipment(db)
	serv := service.Equipment{Repository: rep}
	r.Get("/equipment", serv.GetAll)
	r.Post("/equipment", serv.Post)
}
