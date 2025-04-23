package data

type Todo struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	State bool   `json:"state"`
}
