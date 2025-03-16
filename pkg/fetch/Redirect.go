//go:build wasm

package fetch

type Redirect string

const (
	RedirectFollow Redirect = "follow"
	RedirectError  Redirect = "error"
	RedirectManual Redirect = "manual"
)

func (redirect Redirect) String() string {
	return string(redirect)
}
