package todo

// UpdateTodoRequest struct
type UpdateTodoRequest struct {
	Task        string `json:",omitempty"`
	IsCompleted bool   `json:",omitempty"`
}
