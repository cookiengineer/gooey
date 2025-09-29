package layout

import "github.com/cookiengineer/gooey/components"

func RegisterTo(document *components.Document) {

	document.Register("article", components.Wrap(ToArticle))
	document.Register("dialog",  components.Wrap(ToDialog))
	document.Register("footer",  components.Wrap(ToFooter))
	document.Register("header",  components.Wrap(ToHeader))

}
