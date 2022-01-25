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

type StateHandler struct {
	stateSrv service.StateService
	binder   binder.InputBinder
}

func NewStateHandler(binder binder.InputBinder, stateSrv service.StateService) *StateHandler {
	return &StateHandler{
		stateSrv: stateSrv,
		binder:   binder,
	}
}

// @Summary  Get all states
// @Security ApiKeyAuth
// @Tags state
// @Accept  json
// @Produce  json
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /state [get]
func (h *StateHandler) getStateList(w http.ResponseWriter, r *http.Request) {
	states, _ := h.stateSrv.GetAll()

	render.Respond(w, r, resp.ItemsResponse{Items: states})
}

type createStateResponse struct {
	Id uint `json:"id"`
}

// @Summary  Create state
// @Security ApiKeyAuth
// @Tags state
// @Accept  json
// @Produce  json
// @Success 200 {object} createStateResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /state [post]
func (h *StateHandler) createState(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	state := new(domain.State)

	if err = json.Unmarshal(data, state); err != nil {
		http.Error(w, http.StatusText(422), 422)
		logrus.Debugf("Couldn't unmarshal: %s", err)
		return
	}

	state, err = h.stateSrv.CreateState(state)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		logrus.Debugf("Couldn't create %#v: %s", state, err)
	}

	render.Respond(w, r, createStateResponse{state.ID})
}
