package controllers

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/udistrital/academica_core_service/services"
)

type EstudiantesController struct {
	beego.Controller
	service *services.EstudiantesService
}

func (c *EstudiantesController) Prepare() {
	c.service = services.NewEstudiantesService()
}

func (c *EstudiantesController) DatosDiploma() {
	codigo, err := strconv.ParseInt(c.Ctx.Input.Param(":codigo"), 10, 64)
	if err != nil || codigo <= 0 {
		writeError(&c.Controller, 400, "codigo de estudiante invalido")
		return
	}

	data, err := c.service.GetDatosDiploma(c.Ctx.Request.Context(), codigo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(&c.Controller, 404, "estudiante no encontrado")
			return
		}
		logs.Error("error consultando datos de diploma del estudiante %d: %v", codigo, err)
		writeError(&c.Controller, 500, "error consultando datos de diploma")
		return
	}

	writeJSON(&c.Controller, 200, data)
}
