package controllers

import beego "github.com/beego/beego/v2/server/web"

type apiError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func writeJSON(c *beego.Controller, status int, payload interface{}) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = payload
	c.ServeJSON()
}

func writeError(c *beego.Controller, status int, message string) {
	writeJSON(c, status, apiError{
		Success: false,
		Message: message,
	})
}
