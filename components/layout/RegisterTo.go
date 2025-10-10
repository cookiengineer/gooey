package layout

import "github.com/cookiengineer/gooey/components"

func RegisterTo(document *components.Document) {

	document.Register("article", components.WrapComponent(ToArticle))
	document.Register("dialog",  components.WrapComponent(ToDialog))
	document.Register("footer",  components.WrapComponent(ToFooter))
	document.Register("header",  components.WrapComponent(ToHeader))

}
