//go:build wasm

package app

import "gooey/fetch"
import "syscall/js"

var client_listener_id uint = 0

type ClientListener struct {
	Id       uint                   `json:"id"`
	Callback ClientListenerCallback `json:"callback"`
	Function *js.Func               `json:"function"`
}

type ClientListenerCallback func(string, *fetch.Response, error)

func ToClientListener(callback ClientListenerCallback) ClientListener {

	var listener ClientListener

	listener.Id = client_listener_id
	listener.Callback = callback

	client_listener_id += 1

	return listener

}
