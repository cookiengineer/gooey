package main

import "example/schemas"
import "encoding/json"
import "fmt"
import "io"
import "log"
import "net/http"
import "os"
import "strconv"

func newTaskId(tasks map[int]*schemas.Task) int {

	var found = 0

	for _, task := range tasks {

		if task.ID > found {
			found = task.ID
		}

	}

	if found > 0 {
		found = found + 1
	}

	return found

}

func main() {

	fsys := os.DirFS("public")
	fsrv := http.FileServer(http.FS(fsys))

	tasks := make(map[int]*schemas.Task)

	tasks[1] = &schemas.Task{
		ID:    1,
		Title: "Check out Gooey",
		Done:  true,
	}

	tasks[3] = &schemas.Task{
		ID:    3,
		Title: "Check out Gooey Examples",
		Done:  false,
	}

	tasks[7] = &schemas.Task{
		ID:    7,
		Title: "Build an App",
		Done:  false,
	}

	http.Handle("/", fsrv)

	http.HandleFunc("/api/tasks", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			payload, err := json.MarshalIndent(schemas.Tasks{
				Tasks: tasks,
			}, "", "\t")

			if err == nil {

				fmt.Println("> GET /api/tasks: ok")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusOK)
				response.Write(payload)

			} else {

				fmt.Println("> GET /api/tasks: error")

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte("[]"))

			}

		} else {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("[]"))

		}

	})

	http.HandleFunc("/api/tasks/{id}", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet {

			id, err1 := strconv.ParseInt(request.PathValue("id"), 10, 64)

			if err1 == nil {

				task, ok := tasks[int(id)]

				if ok == true {

					payload, err2 := json.MarshalIndent(task, "", "\t")

					if err2 == nil {

						fmt.Println("> GET /api/tasks/" + strconv.Itoa(int(id)) + ": ok")

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusOK)
						response.Write(payload)

					} else {

						fmt.Println("> GET /api/tasks/" + strconv.Itoa(int(id)) + ": error")

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusInternalServerError)
						response.Write([]byte("{}"))

					}

				} else {

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusNotFound)
					response.Write([]byte("{}"))

				}

			} else {

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("{}"))

			}

		} else if request.Method == http.MethodPost && request.PathValue("id") == "0" {

			bytes, err0 := io.ReadAll(request.Body)

			if err0 == nil {

				schema := schemas.Task{}
				err1 := json.Unmarshal(bytes, &schema)

				if err1 == nil {

					id := newTaskId(tasks)

					if id != 0 {

						schema.ID = id
						tasks[schema.ID] = &schema

						payload, err2 := json.MarshalIndent(schema, "", "\t")

						if err2 == nil {

							fmt.Println("> POST /api/tasks/" + strconv.Itoa(int(id)) + ": ok")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusOK)
							response.Write(payload)

						} else {

							fmt.Println("> POST /api/tasks/" + strconv.Itoa(int(id)) + ": error")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("{}"))

						}

					} else {

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusInternalServerError)
						response.Write([]byte("{}"))

					}

				} else {

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte("{}"))

				}

			} else {

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("{}"))

			}

		} else if request.Method == http.MethodPatch && request.PathValue("id") != "0" {

			bytes, err0 := io.ReadAll(request.Body)

			if err0 == nil {

				schema := schemas.Task{}
				err1 := json.Unmarshal(bytes, &schema)

				if err1 == nil {

					task, ok := tasks[schema.ID]

					if ok == true {

						task.Title = schema.Title
						task.Done = schema.Done

						payload, err2 := json.MarshalIndent(task, "", "\t")

						if err2 == nil {

							fmt.Println("> PATCH /api/tasks/" + strconv.Itoa(schema.ID) + ": ok")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusOK)
							response.Write(payload)

						} else {

							fmt.Println("> PATCH /api/tasks/" + strconv.Itoa(schema.ID) + ": error")

							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusInternalServerError)
							response.Write([]byte("{}"))

						}

					} else {

						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusNotFound)
						response.Write([]byte("{}"))

					}

				} else {

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusBadRequest)
					response.Write([]byte("{}"))

				}

			} else {

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusBadRequest)
				response.Write([]byte("{}"))

			}

		} else {

			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusMethodNotAllowed)
			response.Write([]byte("{}"))

		}

	})

	fmt.Println("Listening on http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
