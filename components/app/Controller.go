//go:build wasm

package app

type Controller struct {
	name   string `json:"name"`
	Main   *Main  `json:"main"`
	Schema any    `json:"schema"`
	View   *View  `json:"view"`
}

func NewController(name string, main *Main, view *View) *Controller {

	var controller Controller

	controller.name = name
	controller.Main = main
	controller.Schema = nil
	controller.View = view

	return &controller

}

func (controller *Controller) Name() string {
	return controller.name
}

func (controller *Controller) SetMain(main *Main) bool {

	controller.Main = main

	return true

}

func (controller *Controller) SetName(name string) bool {

	controller.name = name

	return true

}

func (controller *Controller) SetSchema(schema any) bool {

	controller.Schema = schema

	return true

}

func (controller *Controller) SetView(view *View) bool {

	controller.View = view

	return true

}
