package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/very-important-unmutable-organization/equipment/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db, err := gorm.Open(postgres.Open("host=localhost"))
	if err != nil {
		panic(fmt.Sprintf("could not connect to database: %s", err))
	}

	r.Get("/equipment", func(w http.ResponseWriter, r *http.Request) {
		equipments := new(domain.Equipment)
		res := db.Find(equipments)
		if res.Error != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		_, _ = w.Write([]byte(fmt.Sprint(equipments)))
	})
}
