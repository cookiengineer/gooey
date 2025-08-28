//go:build wasm

package app

type Controller struct {
	Name   string `json:"name"`
	Main   *Main  `json:"main"`
	Schema any    `json:"schema"`
	View   *View  `json:"view"`
}

func NewController(name string, main *Main, view *View) *Controller {

	var controller Controller

	controller.Name   = name
	controller.Main   = main
	controller.Schema = nil
	controller.View   = view

	return &controller

}

func (controller *Controller) GetProperty(name string) string {

	var result string

	switch name {
	case "Name":
		result = controller.Name
	}

	return result

}

func (controller *Controller) SetProperty(name string, value string) bool {

	var result bool

	switch name {
	case "Name":
		controller.Name = value
		result          = true
	}

	return result

}

func (controller *Controller) SetSchema(schema any) bool {

	controller.Schema = schema

	return true

}

