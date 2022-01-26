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

type OriginHandler struct {
	originSrv service.OriginService
	binder    binder.InputBinder
}

func NewOriginHandler(binder binder.InputBinder, originSrv service.OriginService) *OriginHandler {
	return &OriginHandler{
		originSrv: originSrv,
		binder:    binder,
	}
}

// @Summary  Get all origins
// @Security ApiKeyAuth
// @Tags origin
// @Accept  json
// @Produce  json
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /origin [get]
func (h *OriginHandler) getOriginList(w http.ResponseWriter, r *http.Request) {
	origins, _ := h.originSrv.GetAll()

	render.Respond(w, r, resp.ItemsResponse{Items: origins})
}

type createOriginResponse struct {
	Id uint `json:"id"`
}

type createOriginRequest struct { // nolint:unused,deadcode
	Type        domain.OriginType
	EmployeeUID string `json:"employee_uid"`
}

// @Summary  Create origin
// @Security ApiKeyAuth
// @Tags origin
// @Accept  json
// @Produce  json
// @Param input body createOriginRequest true "Origin parameters"
// @Success 200 {object} createOriginResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /origin [post]
func (h *OriginHandler) createOrigin(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	origin := new(domain.Origin)

	if err = json.Unmarshal(data, origin); err != nil {
		http.Error(w, http.StatusText(422), 422)
		logrus.Debugf("Couldn't unmarshal: %s", err)
		return
	}

	origin, err = h.originSrv.CreateOrigin(origin)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		logrus.Debugf("Couldn't create %#v: %s", origin, err)
	}

	render.Respond(w, r, createOriginResponse{origin.ID})
}
