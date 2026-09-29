// Package reactive provides the pure-Go, js-free core used by app.Storage and
// app.Scheduler.
//
// Keeping this logic free of js/wasm dependencies allows it to be unit tested
// natively, while the app package layers persistence and animation frame
// scheduling on top.
package reactive
