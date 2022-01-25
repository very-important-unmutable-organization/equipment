package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/service"
	"github.com/very-important-unmutable-organization/equipment/pkg/binder"
	resp "github.com/very-important-unmutable-organization/equipment/pkg/responses"
)

type ItemTypeHandler struct {
	itemTypeSrv service.ItemTypeService
	binder      binder.InputBinder
}

func NewItemTypeHandler(binder binder.InputBinder, itemTypeSrv service.ItemTypeService) *ItemTypeHandler {
	return &ItemTypeHandler{
		itemTypeSrv: itemTypeSrv,
		binder:      binder,
	}
}

// @Summary  Get all item types
// @Security ApiKeyAuth
// @Tags itemType
// @Accept  json
// @Produce  json
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /type [get]
func (h *ItemTypeHandler) getItemTypeList(w http.ResponseWriter, r *http.Request) {
	itemTypes, _ := h.itemTypeSrv.GetAll()

	render.Respond(w, r, resp.ItemsResponse{Items: itemTypes})
}

type createItemTypeResponse struct {
	Id uint `json:"id"`
}

type createItemRequest struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

// @Summary  Create item type
// @Security ApiKeyAuth
// @Tags itemType
// @Accept  json
// @Produce  json
// @Param input body createItemRequest true "User registration data"
// @Success 200 {object} createItemTypeResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /type [post]
func (h *ItemTypeHandler) createItemType(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	itemType := new(domain.ItemType)

	if err = json.Unmarshal(data, itemType); err != nil {
		http.Error(w, http.StatusText(422), 422)
		logrus.Debugf("Couldn't unmarshal: %s", err)
		return
	}

	itemType, err = h.itemTypeSrv.CreateItemType(itemType)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		logrus.Debugf("Couldn't create %#v: %s", itemType, err)
	}

	render.Respond(w, r, createItemTypeResponse{itemType.ID})
}
