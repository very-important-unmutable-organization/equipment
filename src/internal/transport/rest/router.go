package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/very-important-unmutable-organization/equipment/docs"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
	"github.com/very-important-unmutable-organization/equipment/internal/service"
	"github.com/very-important-unmutable-organization/equipment/pkg/binder"
	mw "github.com/very-important-unmutable-organization/equipment/pkg/middleware"
	"net/http"
	_ "time"
)

type RouterConfig struct {
	ApiToken       string
	ApiTokenHeader string
}

type Router struct {
	repos    *repository.Repositories
	services *service.Services
	binder   binder.InputBinder
}

func NewRouter(repos *repository.Repositories, services *service.Services) *Router {
	return &Router{
		repos:    repos,
		services: services,
		binder:   *binder.NewInputBinder(),
	}
}

func (r *Router) Init() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mw.Recoverer)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			return
		}
	})

	router.Get("/swagger/*", httpSwagger.Handler())

	router.Route("/api/v1", func(router chi.Router) {
		router.Mount("/equipment", r.registerEquipmentRouter())
	})

	return router
}

func (r *Router) registerEquipmentRouter() chi.Router {
	h := NewEquipmentHandler(r.binder, *r.services.EquipmentService)
	router := chi.NewRouter()

	router.Get("/", h.getEquipmentList)
	router.Post("/", h.createEquipment)

	return router
}
