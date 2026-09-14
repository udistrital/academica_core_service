package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"github.com/udistrital/academica_core_service/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/estudiantes",
			beego.NSRouter("/datos-diploma/:codigo", &controllers.EstudiantesController{}, "get:DatosDiploma"),
		),
	)
	beego.AddNamespace(ns)
}
