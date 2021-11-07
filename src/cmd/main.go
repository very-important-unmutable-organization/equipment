package main

import (
	"encoding/json"
	"fmt"
	"github.com/very-important-unmutable-organization/equipment/config"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/very-important-unmutable-organization/equipment/internal/domain"
)

func main() {
	cfg := config.Init()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Database,
	)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(fmt.Sprintf("could not connect to database: %s", err))
	}

	if err = db.AutoMigrate(&domain.Equipment{}); err != nil {
		panic(err)
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
