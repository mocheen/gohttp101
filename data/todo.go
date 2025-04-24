package data

type Todo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}
