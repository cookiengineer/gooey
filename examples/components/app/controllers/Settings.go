package controllers

import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/interfaces"
import "example/schemas"

type Settings struct {
	Main   *app.Main         `json:"main"`
	Schema *schemas.Settings `json:"schema"`
	View   *app.View         `json:"view"`
}

func NewSettings(main *app.Main, view interfaces.View) *Settings {

	var controller Settings

	controller.Main   = main
	controller.Schema = &schemas.Settings{}
	controller.View   = view.(*app.View)

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
	controller.View.Render()
}
