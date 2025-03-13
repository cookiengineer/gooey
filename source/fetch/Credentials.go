//go:build wasm

package fetch

type Credentials string

const (
	CredentialsOmit       Credentials = "omit"
	CredentialsSameOrigin Credentials = "same-origin"
	CredentialsInclude    Credentials = "include"
)

func (credentials Credentials) String() string {
	return string(credentials)
}
