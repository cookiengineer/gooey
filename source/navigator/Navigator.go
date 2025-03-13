//go:build wasm

package navigator

import "syscall/js"

var Navigator navigator

type navigator struct {
	AppCodeName         string    `json:"appCodeName"`
	AppName             string    `json:"appName"`
	AppVersion          string    `json:"appVersion"`
	CookieEnabled       bool      `json:"cookieEnabled"`
	DoNotTrack          bool      `json:"doNotTrack"`
	HardwareConcurrency uint      `json:"hardwareConcurrency"`
	Language            string    `json:"language"`
	Languages           []string  `json:"languages"`
	MaxTouchPoints      uint      `json:"maxTouchPoints"`
	OnLine              bool      `json:"onLine"`
	Platform            string    `json:"platform"`
	Product             string    `json:"product"`
	ProductSub          string    `json:"productSub"`
	PDFViewerEnabled    bool      `json:"pdfViewerEnabled"`
	UserAgent           string    `json:"userAgent"`
	Vendor              string    `json:"vendor"`
	VendorSub           string    `json:"vendorSub"`
	Webdriver           bool      `json:"webdriver"`
	Value               *js.Value `json:"value"`

	// TODO: BuildID string `json:"buildID"` // TODO: Not in Chromium?
	// TODO: oscpu   string `json:"oscpu"`   // TODO: Not in Chromium?

}

func init() {

	navigator_value := js.Global().Get("navigator")

	do_not_track := false
	languages := make([]string, 0)

	tmp1 := navigator_value.Get("doNotTrack")

	if !tmp1.IsNull() && !tmp1.IsUndefined() && tmp1.String() == "1" {
		do_not_track = true
	}

	tmp2 := navigator_value.Get("languages")

	if !tmp2.IsNull() && !tmp2.IsUndefined() {

		for t := 0; t < tmp2.Length(); t++ {

			val := tmp2.Index(t)

			if !val.IsNull() && !val.IsUndefined() && val.String() != "" {
				languages = append(languages, val.String())
			}

		}

	}

	Navigator = navigator{
		AppCodeName:         navigator_value.Get("appCodeName").String(),
		AppName:             navigator_value.Get("appName").String(),
		AppVersion:          navigator_value.Get("appVersion").String(),
		CookieEnabled:       navigator_value.Get("cookieEnabled").Bool(),
		DoNotTrack:          do_not_track,
		HardwareConcurrency: uint(navigator_value.Get("hardwareConcurrency").Int()),
		Language:            navigator_value.Get("language").String(),
		Languages:           languages,
		MaxTouchPoints:      uint(navigator_value.Get("maxTouchPoints").Int()),
		OnLine:              navigator_value.Get("onLine").Bool(),
		Platform:            navigator_value.Get("platform").String(),
		Product:             navigator_value.Get("product").String(),
		ProductSub:          navigator_value.Get("productSub").String(),
		PDFViewerEnabled:    navigator_value.Get("pdfViewerEnabled").Bool(),
		UserAgent:           navigator_value.Get("userAgent").String(),
		Vendor:              navigator_value.Get("vendor").String(),
		VendorSub:           navigator_value.Get("vendorSub").String(),
		Webdriver:           navigator_value.Get("webdriver").Bool(),
	}

}
