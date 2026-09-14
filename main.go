package main

import (
	"os"

	_ "github.com/sijms/go-ora/v2"
	_ "github.com/udistrital/academica_core_service/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/joho/godotenv"

	apistatus "github.com/udistrital/utils_oas/v2/apiStatusLib"
	"github.com/udistrital/utils_oas/v2/auditoria"
	"github.com/udistrital/utils_oas/v2/customerror"
	"github.com/udistrital/utils_oas/v2/database"
	"github.com/udistrital/utils_oas/v2/security"
	"github.com/udistrital/utils_oas/v2/xray"

	"github.com/udistrital/academica_core_service/models"
)

func main() {
	_ = godotenv.Load()

	loadDatabaseConfigFromEnv()

	conn, err := database.BuildOracleConnectionString()
	if err != nil {
		logs.Error("error consultando la cadena de conexión: %v", err)
		return
	}

	if err = models.InitDB(conn); err != nil {
		logs.Error("error al conectarse a la base de datos: %v", err)
		return
	}

	allowedOrigins := []string{"*.udistrital.edu.co"}

	if beego.BConfig.RunMode == beego.DEV {
		allowedOrigins = []string{"*"}

		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{
			"GET",
			"PUT",
			"POST",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"User-Agent",
			"X-Amzn-Trace-Id",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
	}))

	apistatus.Init()
	auditoria.InitMiddleware()
	security.SetSecurityHeaders()
	xray.Init()

	beego.ErrorController(&customerror.CustomErrorController{})

	beego.Run()
}

func loadDatabaseConfigFromEnv() {
	if runMode := os.Getenv("ACADEMICA_CORE_SERVICE_RUN_MODE"); runMode != "" {
		beego.BConfig.RunMode = runMode

		if err := beego.AppConfig.Set("runmode", runMode); err != nil {
			logs.Warn(
				"no se pudo configurar runmode desde ACADEMICA_CORE_SERVICE_RUN_MODE: %v",
				err,
			)
		}
	}

	envConfig := map[string]string{
		"parameterStore": "PARAMETER_STORE",
		"ORuser":         "ACADEMICA_CORE_SERVICE_ORUSER",
		"ORpass":         "ACADEMICA_CORE_SERVICE_ORPASS",
		"ORhost":         "ACADEMICA_CORE_SERVICE_ORHOST",
		"ORport":         "ACADEMICA_CORE_SERVICE_ORPORT",
		"ORservice":      "ACADEMICA_CORE_SERVICE_ORSERVICE",
		"ORschema":       "ACADEMICA_CORE_SERVICE_ORSCHEMA",
	}

	for configKey, envKey := range envConfig {
		if value := os.Getenv(envKey); value != "" {
			if err := beego.AppConfig.Set(configKey, value); err != nil {
				logs.Warn(
					"no se pudo configurar %s desde %s: %v",
					configKey,
					envKey,
					err,
				)
			}
		}
	}
}
