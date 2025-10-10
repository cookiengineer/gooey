package actions

import "github.com/cookiengineer/gooey/components/app"
import "example/schemas"
import "encoding/json"
import "strconv"

func UpdateTask(client *app.Client, task *schemas.Task) (*schemas.Task, error) {

	var result_schema *schemas.Task = nil
	var result_error error = nil

	bytes, err0 := json.MarshalIndent(task, "", "\t")

	if err0 == nil {

		response, err1 := client.Update("/api/tasks/"+strconv.Itoa(task.ID), bytes)

		if err1 == nil {

			schema := schemas.Task{}
			err2 := json.Unmarshal(response.Body, &schema)

			if err2 == nil {
				result_schema = &schema
				result_error = nil
			} else {
				result_error = err2
			}

		} else {
			result_error = err1
		}

	} else {
		result_error = err0
	}

	return result_schema, result_error

}
