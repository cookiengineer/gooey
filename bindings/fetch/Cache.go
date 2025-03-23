//go:build wasm

package fetch

type Cache string

const (
	CacheDefault      Cache = "default"
	CacheNoStore      Cache = "no-store"
	CacheReload       Cache = "reload"
	CacheNone         Cache = "no-cache"
	CacheForce        Cache = "force-cache"
	CacheOnlyIfCached Cache = "only-if-cached"
)

func (cache Cache) String() string {
	return string(cache)
}
