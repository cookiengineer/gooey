package controllers

import "example/schemas"
import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/components/app"

type Settings struct {
	Main   *app.Main         `json:"main"`
	Schema *schemas.Settings `json:"schema"`
	View   *app.View         `json:"view"`
}

func NewSettings(main *app.Main) Settings {

	var controller Settings

	element := bindings.Document.QuerySelector("section[data-name=\"settings\"]")
	view    := app.ToView(element, "Settings", "/settings.html")

	controller.Main   = main
	controller.Schema = &schemas.Settings{}
	controller.View   = &view

	return controller

}
