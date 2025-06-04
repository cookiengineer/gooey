//go:build wasm

package canvas2d

import "github.com/cookiengineer/gooey/bindings/console"
import "errors"
import "syscall/js"

type image_state struct {
	bytes *[]byte
	err   error
}

type Image struct {
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Alt    string    `json:"alt"`
	Src    string    `json:"src"`
	Bytes  *[]byte   `json:"bytes"`
	Value  *js.Value `json:"value"`
}

func NewImage(width int, height int, url string) Image {

	var image Image

	bytes := make([]byte, 0)

	image.Width = width
	image.Height = height
	image.Alt = ""
	image.Src = ""
	image.Bytes = &bytes
	image.Value = nil

	image.SetSrc(url)

	return image

}

func ToImage(value js.Value) *Image {

	var image Image

	image.Width = int(value.Get("width").Int())
	image.Height = int(value.Get("height").Int())
	image.Alt = string(value.Get("alt").String())
	image.Src = string(value.Get("src").String())
	image.Value = &value

	return &image

}

func (image *Image) SetSrc(url string) {

	channel := make(chan *image_state)

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		buffer := make([]byte, 0)
		// TODO: Decode image buffer?

		channel <- &image_state{
			bytes: &buffer,
			err:   nil,
		}

		return nil

	})

	defer on_success.Release()

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		channel <- &image_state{
			bytes: nil,
			err:   errors.New("Image: Failed to load URL \"" + url + "\""),
		}

		return nil

	})

	defer on_failure.Release()

	wrapped_image := js.Global().Get("Image").New(image.Width, image.Height)
	wrapped_image.Set("src", url)
	wrapped_image.Set("onload", on_success)
	wrapped_image.Set("onerror", on_failure)

	state := <-channel

	if state.err != nil {

		image.Bytes = nil
		image.Src = url
		image.Value = &wrapped_image

		console.Error(state.err)

	} else {

		image.Bytes = state.bytes
		image.Src = url
		image.Value = &wrapped_image

	}

}
