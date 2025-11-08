//go:build wasm

// Package animations provides integrations with the global functions RequestAnimationFrame and
// CancelAnimationFrame.
//
// Example usage:
//
//	import "github.com/cookiengineer/bindings/animations"
//	import "github.com/cookiengineer/bindings/console"
//
//	console := console.GetConsole()
//
//	// Request an animation frame
//	identifier := animations.RequestAnimationFrame(func(timestamp float64) {
//		console.Log(identifier, timestamp)
//	})
//
//	// Cancel an animation frame
//	animations.CancelAnimationFrame(identifier)
package animations
