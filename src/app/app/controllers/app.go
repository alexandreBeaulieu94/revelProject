package controllers

import (
	"github.com/alexandreBeaulieu94/revelProject"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}
