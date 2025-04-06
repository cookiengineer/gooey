package controllers

import "example/schemas"
// import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"

type Settings struct {
	Main   *app.Main         `json:"main"`
	Schema *schemas.Settings `json:"schema"`
	View   *app.View         `json:"view"`
}

func NewSettings(main *app.Main) Settings {

	var controller Settings

	view := app.NewView("settings", "Settings", "/settings.html")

	controller.Main   = main
	controller.Schema = &schemas.Settings{}
	controller.View   = &view

	return controller

}
