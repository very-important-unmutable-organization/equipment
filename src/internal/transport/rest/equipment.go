package rest

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/service"
	"github.com/very-important-unmutable-organization/equipment/pkg/binder"
	resp "github.com/very-important-unmutable-organization/equipment/pkg/responses"
)

type EquipmentHandler struct {
	equipmentSrv service.EquipmentService
	binder       binder.InputBinder
}

func NewEquipmentHandler(binder binder.InputBinder, equipmentSrv service.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{
		equipmentSrv: equipmentSrv,
		binder:       binder,
	}
}

////@Success 200 {object} responses.ItemsResponse{items=[]domain.Equipment}

// @Summary  Get all equipment
// @Security ApiKeyAuth
// @Tags equipment
// @Accept  json
// @Produce  json
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /equipment [get]
func (h *EquipmentHandler) getEquipmentList(w http.ResponseWriter, r *http.Request) {
	equipments, _ := h.equipmentSrv.GetAll()

	render.Respond(w, r, resp.ItemsResponse{Items: equipments})
}

type createEquipmentResponse struct {
	Id uint `json:"id"`
}

// @Summary  Create equipment
// @Security ApiKeyAuth
// @Tags equipment
// @Accept  json
// @Produce  json
// @Success 200 {object} createEquipmentResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /equipment [post]
func (h *EquipmentHandler) createEquipment(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	equipment := new(domain.Equipment)

	if err = json.Unmarshal(data, equipment); err != nil {
		http.Error(w, http.StatusText(422), 422)
		logrus.Debugf("Couldn't unmarshal: %s", err)
		return
	}

	equipment, err = h.equipmentSrv.CreateEquipment(equipment)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		logrus.Debugf("Couldn't create %#v: %s", equipment, err)
	}

	render.Respond(w, r, createEquipmentResponse{equipment.ID})
}

//TODO: don't know how to mark success response
////@Success 200 {object} createEquipmentResponse

// @Summary  Get equipment by its id
// @Security ApiKeyAuth
// @Tags equipment
// @Accept  json
// @Produce  json
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /equipment/{id} [get]
func (h *EquipmentHandler) getEquipmentById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Respond(w, r, resp.ErrorNotFound(nil))
		return
	}
	equipment, err := h.equipmentSrv.GetById(id)
	if err == gorm.ErrRecordNotFound {
		render.Respond(w, r, resp.ItemsResponse{Items: []domain.Equipment{}})
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	render.Respond(w, r, resp.ItemsResponse{Items: []domain.Equipment{*equipment}})
}
