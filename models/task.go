package models

type Todo struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Done    bool   `json:"done"`
	Created string `json:created`
	Updated string `json:updated`
}

type RequestParamsInsertTodo struct {
	Name string `json:"name"`
}

type Todos []Todo
