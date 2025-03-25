package types

type Layout string

const (
	LayoutGrid Layout = "grid"
	LayoutFlex Layout = "flex"
	LayoutFlow Layout = "flow"
)

func (typ Layout) String() string {
	return string(typ)
}
