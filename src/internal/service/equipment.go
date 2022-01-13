package service

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
	"io"
	"net/http"
)

type Equipment struct {
	Repository repository.Equipment
}

func (e *Equipment) GetAll(w http.ResponseWriter, r *http.Request) {
	equipments, err := e.Repository.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(equipments); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func (e *Equipment) Post(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	eqt := new(domain.Equipment)
	if err = json.Unmarshal(data, eqt); err != nil {
		http.Error(w, http.StatusText(422), 422)
		logrus.Debugf("Couldn't unmarshal: %s", err)
		return
	}
	if err = e.Repository.Post(eqt); err != nil {
		http.Error(w, http.StatusText(500), 500)
		logrus.Debugf("Couldn't create %#v: %s", eqt, err)
	}
}
