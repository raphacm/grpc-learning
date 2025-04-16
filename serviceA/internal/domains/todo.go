package domains

type ToDoRequest struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type ToDoDomain struct {
	ID         string `json:"id"`
	Task       string `json:"task"`
	IsDone     bool   `json:"is_done"`
	Created_at int64  `json:"created_at"`
}
