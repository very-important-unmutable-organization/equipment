package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/very-important-unmutable-organization/equipment/internal/service"
	"github.com/very-important-unmutable-organization/equipment/internal/transport/rest"
	"github.com/very-important-unmutable-organization/equipment/pkg/logger"

	"github.com/very-important-unmutable-organization/equipment/config"
	"github.com/very-important-unmutable-organization/equipment/internal/repository"
	"github.com/very-important-unmutable-organization/equipment/internal/tables"

	_ "github.com/GoAdminGroup/go-admin/adapter/chi"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"

	"github.com/GoAdminGroup/go-admin/engine"
	adminConfig "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	logrus.Println(cfg)
	r := chi.NewRouter()
	eng := engine.Default()

	adminConf := adminConfig.Config{
		Env: adminConfig.EnvLocal,
		Databases: adminConfig.DatabaseList{
			"default": {
				Host:       cfg.Database.Host,
				Port:       cfg.Database.Port,
				User:       cfg.Database.User,
				Pwd:        cfg.Database.Password,
				Name:       cfg.Database.Database,
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     adminConfig.DriverPostgresql,
			},
		},
		UrlPrefix: "admin",
		Store: adminConfig.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}

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

	r := rest.NewRouter(repos, services)

	//eng.HTML("GET", "/admin", datamodel.GetContent)
	err = eng.AddConfig(&adminConf).
		AddGenerators(tables.Generators).
		AddDisplayFilterXssJsFilter().
		Use(r)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic(fmt.Errorf("PORT env isn't set"))
	}
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), r.Init())
}
