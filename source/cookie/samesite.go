package cookie

type SameSite string

const (
	Strict SameSite = "strict"
	Lax    SameSite = "lax"
	None   SameSite = "none"
)

func (s SameSite) ToString() string {
	return string(s)
}
