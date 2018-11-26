package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) HealthCheck() revel.Result {
	return c.RenderText("Ok")
}
