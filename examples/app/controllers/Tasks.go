package controllers

import "example/actions"
import "example/schemas"
// import "github.com/cookiengineer/gooey/bindings"
import "github.com/cookiengineer/gooey/bindings/console"
// import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
// import "sort"
// import "strconv"

type Tasks struct {
	Main   *app.Main      `json:"main"`
	Schema *schemas.Tasks `json:"schema"`
	View   *app.View      `json:"view"`
}

func NewTasks(main *app.Main) Tasks {

	var controller Tasks

	view := app.NewView("tasks", "Tasks", "/index.html")

	controller.Main   = main
	controller.Schema = &schemas.Tasks{}
	controller.View   = &view

	controller.Main.Footer.Component.AddEventListener("action", components.ToComponentListener(func(event string, attributes map[string]string) {

		action, ok := attributes["action"]

		if ok == true {

			if action == "create" {
				controller.Main.Dialog.Show()
			}

		}

	}, false))

	controller.Main.Dialog.Component.AddEventListener("action", components.ToComponentListener(func(event string, attributes map[string]string) {

		action, ok := attributes["action"]

		if ok == true {

			if action == "confirm" {

				if len(controller.Main.Dialog.Content) > 0 {

					// TODO: How to have dialogs with content.Fieldset?
					// TODO: Maybe something like content.ToFieldset() is necessary?

					fieldset, ok := controller.Main.Dialog.Content[0].(content.Fieldset)

					if ok == true {

						title := fieldset.ValueOf("title").String()
						done  := fieldset.ValueOf("done").Bool()

						task := schemas.Task{
							ID:    0,
							Title: title,
							Done:  done,
						}

						if task.Title != "" {

							controller.Main.Dialog.Disable()

							go func() {

								actions.CreateTask(controller.Main.Client, &task)

								controller.Main.Dialog.Enable()
								controller.Main.Dialog.Hide()

							}()

						}

					} else {
						console.Error("PANIC: No Fieldset!")
					}

				}

				// TODO: Create Task
				console.Log("Create Task Now!")

			} else if action == "cancel" {

				controller.Main.Dialog.Hide()

			}

		}

	}, false))

	console.Log(controller)

	return controller

}

// func (view Tasks) Properties() (string, string, string) {
// 	return "tasks", "Tasks", "/index.html"
// }
// 
// func (view Tasks) BindEvents() {
// 
// 	table  := view.GetElement("table")
// 	dialog := view.GetElement("dialog")
// 	footer := view.GetElement("footer")
// 
// 	if table != nil {
// 
// 		table.AddEventListener("click", dom.ToEventListener(func(event dom.Event) {
// 
// 			target := event.Target
// 
// 			if target.TagName == "INPUT" && target.GetAttribute("type") == "checkbox" {
// 
// 				row      := target.QueryParent("tr")
// 				num, err := strconv.ParseInt(row.GetAttribute("data-id"), 10, 64)
// 
// 				if err == nil {
// 
// 					id       := int(num)
// 					task, ok := view.Schema.Tasks[id]
// 
// 					if ok == true {
// 
// 						if task.Done == true {
// 							task.Done = false
// 						} else {
// 							task.Done = true
// 						}
// 
// 						go func() {
// 
// 							actions.UpdateTask(view.Main.Client, task)
// 							view.Refresh()
// 
// 						}()
// 
// 					}
// 
// 				}
// 
// 			}
// 
// 		}))
// 
// 	}
// 
// 	if dialog != nil {
// 
// 		dialog.QuerySelector("button[data-action=\"confirm\"]").AddEventListener("click", dom.ToEventListener(func(event dom.Event) {
// 
// 			title := dialog.QuerySelector("input[data-name=\"title\"]").Value.Get("value").String()
// 			done  := dialog.QuerySelector("input[data-name=\"done\"]").Value.Get("checked").Bool()
// 
// 			task := schemas.Task{
// 				ID: 0, // set by backend
// 				Title: title,
// 				Done:  done,
// 			}
// 
// 			if task.Title != "" {
// 
// 				buttons := dialog.QuerySelectorAll("button")
// 
// 				for _, button := range buttons {
// 					button.SetAttribute("disabled", "")
// 				}
// 
// 				go func() {
// 
// 					actions.CreateTask(view.Main.Client, &task)
// 					view.CloseDialog()
// 					view.Refresh()
// 
// 					for _, button := range buttons {
// 						button.RemoveAttribute("disabled")
// 					}
// 
// 				}()
// 
// 			}
// 
// 		}))
// 
// 	}
// 
// 	if footer != nil {
// 
// 		footer.QuerySelector("button[data-action=\"create\"]").AddEventListener("click", dom.ToEventListener(func(event dom.Event) {
// 			dialog.SetAttribute("open", "")
// 		}))
// 
// 	}
// 
// }
// 
// func (view Tasks) Enter() bool {
// 
// 	view.Refresh()
// 
// 	return true
// 
// }
// 
// func (view Tasks) Leave() bool {
// 	return true
// }
// 
// func (view Tasks) Refresh() {
// 
// 	schema, err := actions.GetTasks(view.Main.Client)
// 
// 	if err == nil {
// 		view.Schema.Tasks = schema.Tasks
// 		view.Main.Storage.Write("tasks", schema)
// 	}
// 
// 	view.Render()
// 
// }
// 
// func (view Tasks) Render() {
// 
// 	table := view.GetElement("table")
// 
// 	if table != nil {
// 
// 		html := ""
// 		ids  := make([]int, 0)
// 
// 		for _, task := range view.Schema.Tasks {
// 			ids = append(ids, task.ID)
// 		}
// 
// 		sort.Ints(ids)
// 
// 		for i := 0; i < len(ids); i++ {
// 
// 			task := view.Schema.Tasks[ids[i]]
// 			html += view.RenderTask(task)
// 
// 		}
// 
// 		tbody := table.QuerySelector("tbody")
// 
// 		if tbody != nil {
// 			tbody.SetInnerHTML(html)
// 		}
// 
// 	}
// 
// }
// 
// func (view Tasks) RenderTask(task *schemas.Task) string {
// 
// 	var result string
// 
// 	id := strconv.Itoa(task.ID)
// 
// 	result += "<tr data-id=\"" + id + "\">"
// 	result += "<td>" + strconv.Itoa(task.ID) + "</td>"
// 	result += "<td>" + task.Title + "</td>"
// 
// 	if task.Done == true {
// 		result += "<td><input type=\"checkbox\" checked /></td>"
// 	} else {
// 		result += "<td><input type=\"checkbox\" /></td>"
// 	}
// 
// 	result += "</tr>"
// 
// 	return result
// 
// }
// 
// func (view Tasks) CloseDialog() {
// 
// 	dialog := view.GetElement("dialog")
// 
// 	if dialog != nil {
// 
// 		texts := dialog.QuerySelectorAll("input[type=\"text\"]")
// 		bools := dialog.QuerySelectorAll("input[type=\"checkbox\"]")
// 
// 		for _, element := range texts {
// 			element.Value.Set("value", "")
// 		}
// 
// 		for _, element := range bools {
// 			element.Value.Set("checked", false)
// 		}
// 
// 		dialog.RemoveAttribute("open")
// 
// 	}
// 
// }
