package actions

import "github.com/cookiengineer/gooey/components/app"
import "example/schemas"
import "encoding/json"

func GetTasks(client *app.Client) (*schemas.Tasks, error) {

	var result_schema *schemas.Tasks = nil
	var result_error error = nil

	response, err1 := client.Read("/api/tasks")

	if err1 == nil {

		schema := schemas.Tasks{}
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

	return result_schema, result_error

}
