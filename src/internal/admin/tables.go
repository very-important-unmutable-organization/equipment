package admin

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/chi"
	"github.com/GoAdminGroup/go-admin/engine"
	adminConfig "github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"github.com/GoAdminGroup/go-admin/modules/language"
	_ "github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/go-chi/chi"

	"github.com/very-important-unmutable-organization/equipment/config"
)

func Init(cfg *config.DatabaseConfig, router *chi.Mux) {
	eng := engine.Default()
	var Generators = map[string]table.Generator{
		"equipment":  GetEquipmentTable,
		"item_types": GetItemTypeTable,
		"purposes":   GetPurposeTable,
	}

	adminConf := adminConfig.Config{
		Env: adminConfig.EnvLocal,
		Databases: adminConfig.DatabaseList{
			"default": {
				Host:       cfg.Host,
				Port:       cfg.Port,
				User:       cfg.User,
				Pwd:        cfg.Password,
				Name:       cfg.Database,
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

	//eng.HTML("GET", "/admin", datamodel.GetContent)
	err := eng.AddConfig(&adminConf).
		AddGenerators(Generators).
		AddDisplayFilterXssJsFilter().
		Use(router)

	if err != nil {
		panic(err)
	}
}
