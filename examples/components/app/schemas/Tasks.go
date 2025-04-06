package schemas

type Tasks struct {
	Tasks map[int]*Task `json:"tasks"`
}
