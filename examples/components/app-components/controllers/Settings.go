package controllers

import "example/schemas"
import "github.com/cookiengineer/gooey/components/app"

type Settings struct {
	Main   *app.Main         `json:"main"`
	Schema *schemas.Settings `json:"schema"`
	View   *app.View         `json:"view"`
}

func NewSettings(main *app.Main, view *app.View) *Settings {

	var controller Settings

	controller.Main   = main
	controller.Schema = &schemas.Settings{}
	controller.View   = view

	// TODO

	return &controller

}

func (controller *Settings) Name() string {
	return "settings"
}

func (controller *Settings) Update() {
	// Not Implemented
}

func (controller *Settings) Render() {
	// Not Implemented
}
