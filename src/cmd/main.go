package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/very-important-unmutable-organization/equipment/config"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"

	"github.com/sirupsen/logrus"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

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
			return
		}
		res := db.Create(eqt)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		panic(fmt.Errorf("PORT env isn't set"))
	}
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
