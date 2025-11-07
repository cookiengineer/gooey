package controllers

import "example/schemas"
import "example/views"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/interfaces"

type Settings struct {
	Main   *app.Main         `json:"main"`
	Schema *schemas.Settings `json:"schema"`
	View   *views.Settings   `json:"view"`
}

func NewSettings(main *app.Main, view interfaces.View) *Settings {

	var controller Settings

	controller.Main   = main
	controller.Schema = &schemas.Settings{}
	controller.View   = view.(*views.Settings)

	return &controller

}

func (controller *Settings) Enter() bool {

	controller.Update()

	// TODO: Add EventListeners

	return true

}

func (controller *Settings) Leave() bool {

	// TODO: Remove EventListeners

	return true

}

func (controller *Settings) Name() string {
	return "settings"
}

func (controller *Settings) Update() {

	// Not Implemented
	controller.Render()

}

func (controller *Settings) Render() {
	controller.View.Render()
}
