package rest

import (
	"net/http"
	_ "time"

	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/GoAdminGroup/go-admin/adapter/chi"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"github.com/go-chi/chi"

	_ "github.com/very-important-unmutable-organization/equipment/docs"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
	"github.com/very-important-unmutable-organization/equipment/internal/service"
	"github.com/very-important-unmutable-organization/equipment/pkg/binder"
	mw "github.com/very-important-unmutable-organization/equipment/pkg/middleware"
)

type RouterConfig struct {
	ApiToken       string
	ApiTokenHeader string
}

type Router struct {
	repos    *repository.Repositories
	services *service.Services
	binder   binder.InputBinder
	config   RouterConfig
}

func NewRouter(repos *repository.Repositories, services *service.Services, cfg RouterConfig) *Router {
	return &Router{
		repos:    repos,
		services: services,
		binder:   *binder.NewInputBinder(),
		config:   cfg,
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
		router.Use(mw.ApiKeyAuthentication(r.config.ApiToken, r.config.ApiTokenHeader))
		router.Mount("/equipment", r.registerEquipmentRouter())
		router.Mount("/type", r.registerItemTypeRouter())
		router.Mount("/state", r.registerStateRouter())
		router.Mount("/purpose", r.registerPurposeRouter())
		router.Mount("/origin", r.registerOriginRouter())
		router.Mount("/document", r.registerDocumentRouter())
		router.Mount("/photo", r.registerPhotoRouter())
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

func (r *Router) registerItemTypeRouter() chi.Router {
	h := NewItemTypeHandler(r.binder, *r.services.ItemTypeService)
	router := chi.NewRouter()

	router.Get("/", h.getItemTypeList)
	router.Post("/", h.createItemType)

	return router
}

func (r *Router) registerStateRouter() chi.Router {
	h := NewStateHandler(r.binder, *r.services.StateService)
	router := chi.NewRouter()

	router.Get("/", h.getStateList)
	router.Post("/", h.createState)

	return router
}

func (r *Router) registerPurposeRouter() chi.Router {
	h := NewPurposeHandler(r.binder, *r.services.PurposeService)
	router := chi.NewRouter()

	router.Get("/", h.getPurposeList)
	router.Post("/", h.createPurpose)

	return router
}

func (r *Router) registerOriginRouter() chi.Router {
	h := NewOriginHandler(r.binder, *r.services.OriginService)
	router := chi.NewRouter()

	router.Get("/", h.getOriginList)
	router.Post("/", h.createOrigin)

	return router
}

func (r *Router) registerDocumentRouter() chi.Router {
	h := NewDocumentHandler(r.binder, *r.services.DocumentService)
	router := chi.NewRouter()

	router.Get("/", h.getDocumentList)
	router.Post("/", h.createDocument)

	return router
}

func (r *Router) registerPhotoRouter() chi.Router {
	h := NewPhotoHandler(r.binder, *r.services.PhotoService)
	router := chi.NewRouter()

	router.Get("/", h.getPhotoList)
	router.Post("/", h.createPhoto)

	return router
}
