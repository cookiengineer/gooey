package layout

import "github.com/cookiengineer/gooey/components"

func RegisterTo(document *components.Document) {

	document.Register("article", Wrap(ToArticle))
	document.Register("dialog",  Wrap(ToDialog))
	document.Register("footer",  Wrap(ToFooter))
	document.Register("header",  Wrap(ToHeader))

}
