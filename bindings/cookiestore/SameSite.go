package cookiestore

type SameSite string

const (
	SameSiteStrict SameSite = "strict"
	SameSiteLax    SameSite = "lax"
	SameSiteNone   SameSite = "none"
)

func (samesite SameSite) String() string {
	return string(samesite)
}
