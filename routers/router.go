package routers

import (
	"github.com/astaxie/beego"

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
