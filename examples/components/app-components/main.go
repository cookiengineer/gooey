package main

// import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/interfaces"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/content"
// import "github.com/cookiengineer/gooey/components/layout"
// import "encoding/json"
import "fmt"
import "time"

func main() {

	document := components.NewDocument()
	document.Register("fieldset", func(element *dom.Element) interfaces.Component {
		return content.ToFieldset(element)
	})

	fmt.Println(document)

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
