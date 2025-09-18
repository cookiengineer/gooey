package controllers

import "example/actions"
import "example/schemas"
import "github.com/cookiengineer/gooey/bindings/console"
import "github.com/cookiengineer/gooey/bindings/dom"
import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/data"
import "github.com/cookiengineer/gooey/components/layout"
import "sync"

type Tasks struct {
	Main   *app.Main      `json:"main"`
	Schema *schemas.Tasks `json:"schema"`
	View   *app.View      `json:"view"`
}

func NewTasks(main *app.Main) Tasks {

	var controller Tasks

	element := dom.Document.QuerySelector("section[data-name=\"tasks\"]")
	view    := app.ToView(element, "Tasks", "/index.html")

	controller.Main   = main
	controller.Schema = &schemas.Tasks{}
	controller.View   = view

	table := controller.queryTable()

	if table != nil {

		table.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

			action, ok := attributes["action"].(string)

			if ok == true {

				if action == "mark-undone" {

					go func() {

						indexes, dataset := table.Selected()
						waitgroup := sync.WaitGroup{}

						table.Deselect(indexes)

						for d := 0; d < len(dataset); d++ {

							index := indexes[d]
							entry := dataset[d]

							waitgroup.Add(1)

							go func(index int, entry data.Data) {

								entry["done"] = false

								task := schemas.Task{
									ID:    entry["id"].(int),
									Title: entry["title"].(string),
									Done:  entry["done"].(bool),
								}

								result, err := actions.UpdateTask(controller.Main.Client, &task)

								if result.Done == false && err == nil {

									table.Remove([]int{index})
									table.Add(entry)

								}

								defer waitgroup.Done()

							}(index, entry)

						}

						waitgroup.Wait()
						table.Render()

					}()

				} else if action == "mark-done" {

					go func() {

						indexes, dataset := table.Selected()
						waitgroup := sync.WaitGroup{}

						table.Deselect(indexes)

						for d := 0; d < len(dataset); d++ {

							index := indexes[d]
							entry := dataset[d]

							waitgroup.Add(1)

							go func(index int, entry data.Data) {

								entry["done"] = true

								task := schemas.Task{
									ID:    entry["id"].(int),
									Title: entry["title"].(string),
									Done:  entry["done"].(bool),
								}

								result, err := actions.UpdateTask(controller.Main.Client, &task)

								if result.Done == true && err == nil {

									table.Remove([]int{index})
									table.Add(entry)

								}

								defer waitgroup.Done()

							}(index, entry)

						}

						waitgroup.Wait()
						table.Render()

					}()

				}

			}

		}, false))

	}

	controller.Main.Footer.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		action, ok := attributes["action"].(string)

		if ok == true {

			if action == "create" {
				controller.Main.Dialog.Show()
			}

		}

	}, false))

	controller.Main.Dialog.Component.AddEventListener("action", components.ToEventListener(func(event string, attributes map[string]any) {

		action, ok := attributes["action"].(string)

		if ok == true {

			if action == "confirm" {

				go func() {

					fieldset := controller.queryFieldset()
					table := controller.queryTable()

					if fieldset != nil && table != nil {

						title := fieldset.ValueOf("title").String()
						done  := fieldset.ValueOf("done").Bool()

						task := schemas.Task{
							ID:    0,
							Title: title,
							Done:  done,
						}

						if task.Title != "" {

							fieldset.Reset()
							controller.Main.Dialog.Disable()

							waitgroup := sync.WaitGroup{}
							waitgroup.Add(1)

							go func() {

								result, err := actions.CreateTask(controller.Main.Client, &task)

								if result.Title == task.Title && err == nil {

									table.Add(data.Data(map[string]any{
										"id":    result.ID,
										"title": task.Title,
										"done":  task.Done,
									}))

								}

								controller.Main.Dialog.Enable()
								controller.Main.Dialog.Hide()

								defer waitgroup.Done()

							}()

							waitgroup.Wait()
							table.Render()

						}

					}

				}()

			} else if action == "cancel" {

				go func() {

					fieldset := controller.queryFieldset()

					if fieldset != nil {

						fieldset.Reset()

						controller.Main.Dialog.Enable()
						controller.Main.Dialog.Hide()

					}

				}()

			}

		}

	}, false))

	controller.Update()

	return controller

}

func (controller *Tasks) Update() {

	if controller.Main != nil {

		schema, err := actions.GetTasks(controller.Main.Client)

		if err == nil {
			controller.Schema = schema
			controller.Main.Storage.Write("tasks", schema)
		}

		controller.Render()

	}

}

func (controller *Tasks) Render() {

	for c := 0; c < len(controller.View.Content); c++ {

		article, ok1 := controller.View.Content[c].(*layout.Article)

		if ok1 == true && len(article.Content) == 1 {

			table, ok2 := article.Content[0].(*content.Table)

			if len(controller.Schema.Tasks) > 0 && ok2 == true {

				dataset := data.NewDataset(0)

				for _, task := range controller.Schema.Tasks {

					dataset.Add(data.Data(map[string]any{
						"id": task.ID,
						"title": task.Title,
						"done": task.Done,
					}))

				}

				console.Log(dataset)
				console.Log(table)

				table.SetDataset(dataset)
				table.SortBy("id")

				table.Render()

			}

		}


	}

}

func (controller *Tasks) queryFieldset() *content.Fieldset {

	var result *content.Fieldset

	if controller.Main.Dialog.Content != nil {

		fieldset, ok1 := controller.Main.Dialog.Content.(*content.Fieldset)

		if ok1 == true {
			result = fieldset
		}

	}

	return result

}

func (controller *Tasks) queryTable() *content.Table {

	var result *content.Table

	if len(controller.View.Content) > 0 {

		article, ok1 := controller.View.Content[0].(*layout.Article)

		if ok1 == true && len(article.Content) > 0 {

			table, ok2 := article.Content[0].(*content.Table)

			if ok2 == true {
				result = table
			}

		}

	}

	return result

}

