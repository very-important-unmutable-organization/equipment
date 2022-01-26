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

type DocumentHandler struct {
	documentSrv service.DocumentService
	binder      binder.InputBinder
}

func NewDocumentHandler(binder binder.InputBinder, documentSrv service.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		documentSrv: documentSrv,
		binder:      binder,
	}
}

// @Summary  Get all documents
// @Security ApiKeyAuth
// @Tags document
// @Accept  json
// @Produce  json
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /document [get]
func (h *DocumentHandler) getDocumentList(w http.ResponseWriter, r *http.Request) {
	documents, _ := h.documentSrv.GetAll()

	render.Respond(w, r, resp.ItemsResponse{Items: documents})
}

type createDocumentResponse struct {
	Id uint `json:"id"`
}

type createDocumentRequest struct { // nolint:unused,deadcode
	ItemID   uint `json:"item_id"`
	Location string
}

// @Summary  Create document
// @Security ApiKeyAuth
// @Tags document
// @Accept  json
// @Produce  json
// @Param input body createDocumentRequest true "Document parameters"
// @Success 200 {object} createDocumentResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Router /document [post]
func (h *DocumentHandler) createDocument(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	document := new(domain.Document)

	if err = json.Unmarshal(data, document); err != nil {
		http.Error(w, http.StatusText(422), 422)
		logrus.Debugf("Couldn't unmarshal: %s", err)
		return
	}

	document, err = h.documentSrv.CreateDocument(document)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		logrus.Debugf("Couldn't create %#v: %s", document, err)
	}

	render.Respond(w, r, createDocumentResponse{document.ID})
}
