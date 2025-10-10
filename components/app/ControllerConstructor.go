//go:build wasm

package app

import "github.com/cookiengineer/gooey/components/interfaces"

type ControllerConstructor func(*Main, interfaces.View) interfaces.Controller
