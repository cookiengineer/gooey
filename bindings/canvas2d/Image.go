//go:build wasm

package canvas2d

import "syscall/js"

type Image struct {
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Alt    string    `json:"alt"`
	Src    string    `json:"src"`
	Value  *js.Value `json:"value"`
}

func NewImage(width int, height int, url string) Image {

	var image Image

	// TODO: Set src property
	// TODO: call await load()

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

	// TODO: onerror
	// TODO: onload
	// TODO: channel for go func()

}
