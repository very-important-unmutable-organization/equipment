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

type PhotoHandler struct {
	photoSrv service.PhotoService
	binder   binder.InputBinder
}

func NewPhotoHandler(binder binder.InputBinder, photoSrv service.PhotoService) *PhotoHandler {
	return &PhotoHandler{
		photoSrv: photoSrv,
		binder:   binder,
	}
}

// @Summary  Get all photos
// @Security ApiKeyAuth
// @Tags photo
// @Accept  json
// @Produce  json
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /photo [get]
func (h *PhotoHandler) getPhotoList(w http.ResponseWriter, r *http.Request) {
	photos, _ := h.photoSrv.GetAll()

	render.Respond(w, r, resp.ItemsResponse{Items: photos})
}

type createPhotoResponse struct {
	Id uint `json:"id"`
}

type createPhotoRequest struct { // nolint:unused,deadcode
	ItemID   uint `json:"item_id"`
	Location string
}

// @Summary  Create photo
// @Security ApiKeyAuth
// @Tags photo
// @Accept  json
// @Produce  json
// @Param input body createPhotoRequest true "Photo parameters"
// @Success 200 {object} createPhotoResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /photo [post]
func (h *PhotoHandler) createPhoto(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	photo := new(domain.Photo)

	if err = json.Unmarshal(data, photo); err != nil {
		http.Error(w, http.StatusText(422), 422)
		logrus.Debugf("Couldn't unmarshal: %s", err)
		return
	}

	photo, err = h.photoSrv.CreatePhoto(photo)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		logrus.Debugf("Couldn't create %#v: %s", photo, err)
	}

	render.Respond(w, r, createPhotoResponse{photo.ID})
}
