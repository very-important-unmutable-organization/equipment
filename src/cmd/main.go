package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/config"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"

	"github.com/sirupsen/logrus"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

// TODO: enums don't work, probably 'cause they weren't created

func main() {
	cfg := config.Init()
	logrus.Println(cfg)

	db, err := repository.InitDb(repository.Config(cfg.Database))
	if err != nil {
		logrus.Errorf("error occured while initialzing db client: %s", err.Error())
		return
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	registerEquipmentEndpoints(r, db)
	registerTypeEndpoints(r, db)
	registerStateEndpoints(r, db)
	registerPurposeEndpoints(r, db)
	registerOriginEndpoints(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		panic(fmt.Errorf("PORT env isn't set"))
	}
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func registerOriginEndpoints(r *chi.Mux, db *gorm.DB) {
	r.Get("/origin", func(w http.ResponseWriter, r *http.Request) {
		items := new([]domain.Origin)
		res := db.Find(items)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		enc := json.NewEncoder(w)
		if err := enc.Encode(items); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	})

	r.Post("/origin", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		item := new(domain.Origin)
		if err = json.Unmarshal(data, item); err != nil {
			http.Error(w, http.StatusText(422), 422)
			logrus.Debugf("Couldn't unmarshal: %s", err)
			return
		}
		res := db.Create(item)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			logrus.Debugf("Couldn't create %#v: %s", item, err)
			return
		}
	})
}

func registerPurposeEndpoints(r *chi.Mux, db *gorm.DB) {
	r.Get("/purpose", func(w http.ResponseWriter, r *http.Request) {
		items := new([]domain.Purpose)
		res := db.Find(items)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		enc := json.NewEncoder(w)
		if err := enc.Encode(items); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	})

	r.Post("/purpose", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		item := new(domain.Purpose)
		if err = json.Unmarshal(data, item); err != nil {
			http.Error(w, http.StatusText(422), 422)
			logrus.Debugf("Couldn't unmarshal: %s", err)
			return
		}
		logrus.Debugf("Adding %#v", item)
		res := db.Create(item)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			logrus.Debugf("Couldn't create %#v: %s", item, err)
			return
		}
	})
}

func registerStateEndpoints(r *chi.Mux, db *gorm.DB) {
	r.Get("/state", func(w http.ResponseWriter, r *http.Request) {
		states := new([]domain.State)
		res := db.Find(states)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		enc := json.NewEncoder(w)
		if err := enc.Encode(states); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	})

	r.Post("/state", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		item := new(domain.State)
		if err = json.Unmarshal(data, item); err != nil {
			http.Error(w, http.StatusText(422), 422)
			logrus.Debugf("Couldn't unmarshal: %s", err)
			return
		}
		res := db.Create(item)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			logrus.Debugf("Couldn't create %#v: %s", item, err)
			return
		}
	})
}

func registerTypeEndpoints(r *chi.Mux, db *gorm.DB) {
	r.Get("/type", func(w http.ResponseWriter, r *http.Request) {
		types := new([]domain.ItemType)
		res := db.Find(types)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		enc := json.NewEncoder(w)
		if err := enc.Encode(types); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	})

	r.Post("/type", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		item := new(domain.ItemType)
		if err = json.Unmarshal(data, item); err != nil {
			http.Error(w, http.StatusText(422), 422)
			logrus.Debugf("Couldn't unmarshal: %s", err)
			return
		}
		res := db.Create(item)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			logrus.Debugf("Couldn't create %#v: %s", item, err)
			return
		}
	})
}

func registerEquipmentEndpoints(r *chi.Mux, db *gorm.DB) {
	r.Get("/equipment", func(w http.ResponseWriter, r *http.Request) {
		equipments := new([]domain.Equipment)
		res := db.Find(equipments)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		enc := json.NewEncoder(w)
		if err := enc.Encode(equipments); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	})

	r.Post("/equipment", func(w http.ResponseWriter, r *http.Request) {
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
		res := db.Create(eqt)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			logrus.Debugf("Couldn't create %#v: %s", eqt, err)
			return
		}
	})
}
