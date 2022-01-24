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

type PurposeHandler struct {
	purposeSrv service.PurposeService
	binder     binder.InputBinder
}

func NewPurposeHandler(binder binder.InputBinder, purposeSrv service.PurposeService) *PurposeHandler {
	return &PurposeHandler{
		purposeSrv: purposeSrv,
		binder:     binder,
	}
}

// @Summary  Get all purposes
// @Security ApiKeyAuth
// @Tags purpose
// @Accept  json
// @Produce  json
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /purpose [get]
func (h *PurposeHandler) getPurposeList(w http.ResponseWriter, r *http.Request) {
	purposes, _ := h.purposeSrv.GetAll()

	render.Respond(w, r, resp.ItemsResponse{Items: purposes})
}

type createPurposeResponse struct {
	Id uint `json:"id"`
}

// @Summary  Create purpose
// @Security ApiKeyAuth
// @Tags purpose
// @Accept  json
// @Produce  json
// @Success 200 {object} createPurposeResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /purpose [post]
func (h *PurposeHandler) createPurpose(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	purpose := new(domain.Purpose)

	if err = json.Unmarshal(data, purpose); err != nil {
		http.Error(w, http.StatusText(422), 422)
		logrus.Debugf("Couldn't unmarshal: %s", err)
		return
	}

	purpose, err = h.purposeSrv.CreatePurpose(purpose)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		logrus.Debugf("Couldn't create %#v: %s", purpose, err)
	}
	render.Respond(w, r, createPurposeResponse{purpose.ID})
}
