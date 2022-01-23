package main

import (
	"fmt"
	"github.com/very-important-unmutable-organization/equipment/internal/admin"
	"net/http"
	"os"

	"github.com/very-important-unmutable-organization/equipment/internal/service"
	"github.com/very-important-unmutable-organization/equipment/internal/transport/rest"
	"github.com/very-important-unmutable-organization/equipment/pkg/logger"

	"github.com/very-important-unmutable-organization/equipment/config"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"

	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

// @title Equipment API
// @version 1.0
// @description mem

// @BasePath /api/v1/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY

// @securityDefinitions.apikey UserTokenAuth
// @in header
// @name Authorization

// Run initializes whole application.
func main() {
	cfg := config.Init()

	logger.InitLogger(logger.Config{ServiceName: "equipment-api", Debug: true})

	db, err := repository.InitDb(repository.Config(cfg.Database))
	if err != nil {
		logrus.Errorf("error occured while initialzing db client: %s", err.Error())
		return
	}

	repos, err := repository.NewRepositories(db)
	if err != nil {
		logrus.Errorf("error occured while initialzing db client: %s", err.Error())
		return
	}

	services, _ := service.NewServices(repos)

	r := rest.NewRouter(repos, services).Init()

	admin.Init(&cfg.Database, r)

	port := os.Getenv("PORT")
	if port == "" {
		panic(fmt.Errorf("PORT env isn't set"))
	}
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
