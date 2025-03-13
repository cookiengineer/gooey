//go:build wasm

package app

import "gooey/fetch"
import "bytes"
import "strings"

type Client struct {
	listeners map[string][]*ClientListener
}

func NewClient() Client {

	var client Client

	client.listeners = make(map[string][]*ClientListener)

	return client

}

func (client *Client) AddListener(path string, listener ClientListener) bool {

	var result bool = false

	_, ok := client.listeners[path]

	if ok == true {
		client.listeners[path] = append(client.listeners[path], &listener)
		result = true
	} else {
		client.listeners[path] = make([]*ClientListener, 0)
		client.listeners[path] = append(client.listeners[path], &listener)
		result = true
	}

	return result

}

func (client *Client) RemoveListener(path string, listener *ClientListener) bool {

	var result bool = false

	if listener != nil {

		listeners, ok := client.listeners[path]

		if ok == true {

			var index int = -1

			for l := 0; l < len(listeners); l++ {

				if listeners[l].Id == listener.Id {
					index = l
					break
				}

			}

			if index != -1 {
				client.listeners[path] = append(client.listeners[path][:index], client.listeners[path][index+1:]...)
				result = true
			}

		}

	} else {

		_, ok := client.listeners[path]

		if ok == true {
			delete(client.listeners, path)
			result = true
		}

	}

	return result

}

func (client *Client) Request(path string) {

	if strings.HasPrefix(path, "/api") {

		go func(){

			response, err := fetch.Fetch(path, &fetch.Request{
				Method:      fetch.MethodGET,
				Mode:        fetch.ModeSameOrigin,
				Cache:       fetch.CacheDefault,
				Credentials: fetch.CredentialsOmit,
				Redirect:    fetch.RedirectError,
				Headers:     map[string]string{
					"Accept": "application/json",
				},
			})

			listeners, ok := client.listeners[path]

			if ok == true {

				if err == nil {

					for l := 0; l < len(listeners); l++ {
						listeners[l].Callback(path, response, nil)
					}

				} else {

					for l := 0; l < len(listeners); l++ {
						listeners[l].Callback(path, nil, err)
					}

				}

			}

		}()

	}

}

func (client *Client) Send(path string, payload []byte) {

	if strings.HasPrefix(path, "/api") {

		go func() {

			response, err := fetch.Fetch(path, &fetch.Request{
				Method:      fetch.MethodPOST,
				Mode:        fetch.ModeSameOrigin,
				Cache:       fetch.CacheDefault,
				Credentials: fetch.CredentialsOmit,
				Redirect:    fetch.RedirectError,
				Headers:     map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/json",
				},
				Body: bytes.NewReader(payload),
			})

			listeners, ok := client.listeners[path]

			if ok == true {

				if err == nil {

					for l := 0; l < len(listeners); l++ {
						listeners[l].Callback(path, response, nil)
					}

				} else {

					for l := 0; l < len(listeners); l++ {
						listeners[l].Callback(path, nil, err)
					}

				}

			}

		}()

	}

}

