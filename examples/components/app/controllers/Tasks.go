package controllers

import "github.com/cookiengineer/gooey/components"
import "github.com/cookiengineer/gooey/components/app"
import "github.com/cookiengineer/gooey/components/content"
import "github.com/cookiengineer/gooey/components/data"
import "github.com/cookiengineer/gooey/components/interfaces"
import "example/actions"
import "example/schemas"
import "sync"

type Tasks struct {
	Main        *app.Main      `json:"main"`
	Schema      *schemas.Tasks `json:"schema"`
	View        *app.View      `json:"view"`
	unsubscribe []func()       `json:"-"`
}

func NewTasks(main *app.Main, view interfaces.View) *Tasks {

	var controller Tasks

	controller.Main = main
	controller.Schema = &schemas.Tasks{}
	controller.View = view.(*app.View)

	return &controller

}

func (controller *Tasks) Enter() bool {

	// IMPORTANT: The Component Query API is self-including

	fieldset, ok0 := components.UnwrapComponent[*content.Fieldset](controller.Main.Dialog.Query("dialog > fieldset"))
	table, ok1 := components.UnwrapComponent[*content.Table](controller.View.Query("section > article > table"))

	if fieldset != nil && table != nil && ok0 == true && ok1 == true {

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

					if fieldset != nil && table != nil {

						title := fieldset.ValueOf("title").String()
						done := fieldset.ValueOf("done").Bool()

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

						}

					}

				}()

			} else if action == "cancel" {

				go func() {

					if fieldset != nil {

						fieldset.Reset()

						controller.Main.Dialog.Enable()
						controller.Main.Dialog.Hide()

					}

				}()

			}

		}

	}, false))

	controller.unsubscribe = append(controller.unsubscribe, controller.Main.Storage.Subscribe("tasks", func(value any) {

		schema, ok := value.(*schemas.Tasks)

		if ok == true {
			controller.Schema = schema
			controller.applyTasks(schema)
		}

	}))

	go controller.Update()

	return true

}

func (controller *Tasks) Leave() bool {

	controller.Main.Footer.Component.RemoveEventListener("action", nil)
	controller.Main.Dialog.Component.RemoveEventListener("action", nil)

	for _, unsubscribe := range controller.unsubscribe {
		unsubscribe()
	}

	controller.unsubscribe = controller.unsubscribe[:0]

	return true

}

func (controller *Tasks) applyTasks(schema *schemas.Tasks) {

	table, ok := components.UnwrapComponent[*content.Table](controller.View.Query("section > article > table"))

	if table != nil && ok == true && schema != nil {

		dataset := data.NewDataset(0)

		for _, task := range schema.Tasks {

			dataset.Add(data.Data(map[string]any{
				"id":    task.ID,
				"title": task.Title,
				"done":  task.Done,
			}))

		}

		table.SetDataset(dataset)
		table.SortBy("id")

	}

}

func (controller *Tasks) Name() string {
	return "tasks"
}

func (controller *Tasks) Update() {

	if controller.Main != nil {

		schema, err := actions.GetTasks(controller.Main.Client)

		if err == nil {
			controller.Main.Storage.Update("tasks", schema)
		}

	}

}

func (controller *Tasks) Render() {
	controller.View.Render()
}
