//go:build wasm

package app

import "github.com/cookiengineer/gooey/bindings/fetch"
import "bytes"
import "errors"
import "strconv"
import "strings"

type client_state struct {
	response *fetch.Response
	err      error
}

type Client struct {
	listeners map[string][]*ClientListener
}

func NewClient() *Client {

	var client Client

	client.listeners = make(map[string][]*ClientListener)

	return &client

}

func (client *Client) AddListener(path string, listener ClientListener) bool {

	var result bool

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

	var result bool

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

func (client *Client) Create(path string, payload []byte) (*fetch.Response, error) {

	if strings.HasPrefix(path, "/api") {

		channel := make(chan *client_state)

		go func() {

			response, err := fetch.Fetch(path, &fetch.Request{
				Method:      fetch.MethodPost,
				Mode:        fetch.ModeSameOrigin,
				Cache:       fetch.CacheDefault,
				Credentials: fetch.CredentialsOmit,
				Redirect:    fetch.RedirectError,
				Headers:     map[string]string{
					"Accept":         "application/json",
					"Content-Type":   "application/json",
					"Content-Length": strconv.Itoa(len(payload)),
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

			channel <- &client_state{
				response: response,
				err:      err,
			}

		}()

		state := <-channel

		return state.response, state.err

	} else {
		return nil, errors.New("Invalid Path: Expected /api/* prefix")
	}

}

func (client *Client) Read(path string) (*fetch.Response, error) {

	if strings.HasPrefix(path, "/api") {

		channel := make(chan *client_state)

		go func() {

			response, err := fetch.Fetch(path, &fetch.Request{
				Method:      fetch.MethodGet,
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

			channel <- &client_state{
				response: response,
				err:      err,
			}

		}()

		state := <-channel

		return state.response, state.err

	} else {
		return nil, errors.New("Invalid Path: Expected /api/* prefix")
	}

}

func (client *Client) Update(path string, payload []byte) (*fetch.Response, error) {

	if strings.HasPrefix(path, "/api") {

		channel := make(chan *client_state)

		go func() {

			response, err := fetch.Fetch(path, &fetch.Request{
				Method:      fetch.MethodPatch,
				Mode:        fetch.ModeSameOrigin,
				Cache:       fetch.CacheDefault,
				Credentials: fetch.CredentialsOmit,
				Redirect:    fetch.RedirectError,
				Headers:     map[string]string{
					"Accept":         "application/json",
					"Content-Type":   "application/json",
					"Content-Length": strconv.Itoa(len(payload)),
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

			channel <- &client_state{
				response: response,
				err:      err,
			}

		}()

		state := <-channel

		return state.response, state.err

	} else {
		return nil, errors.New("Invalid Path: Expected /api/* prefix")
	}

}

func (client *Client) Delete(path string, payload []byte) (*fetch.Response, error) {

	if strings.HasPrefix(path, "/api") {

		channel := make(chan *client_state)

		go func() {

			response, err := fetch.Fetch(path, &fetch.Request{
				Method:      fetch.MethodDelete,
				Mode:        fetch.ModeSameOrigin,
				Cache:       fetch.CacheDefault,
				Credentials: fetch.CredentialsOmit,
				Redirect:    fetch.RedirectError,
				Headers:     map[string]string{
					"Accept":         "application/json",
					"Content-Type":   "application/json",
					"Content-Length": strconv.Itoa(len(payload)),
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

			channel <- &client_state{
				response: response,
				err:      err,
			}

		}()

		state := <-channel

		return state.response, state.err

	} else {
		return nil, errors.New("Invalid Path: Expected /api/* prefix")
	}

}

