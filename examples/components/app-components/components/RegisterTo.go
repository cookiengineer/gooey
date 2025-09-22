package components

import "github.com/cookiengineer/gooey/components"

func RegisterTo(document *components.Document) {

	document.Register("app-example", Wrap(ToExample))

}
