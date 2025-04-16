package domains

import (
	"time"

	"github.com/google/uuid"
)

type ToDoRequest struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func (domain *ToDoRequest) ToToDoDomain() ToDoDomain {

	return ToDoDomain{
		ID:         uuid.NewString(),
		Task:       domain.Task,
		IsDone:     domain.IsDone,
		Created_at: time.Now().UnixMilli(),
	}
}

type ToDoDomain struct {
	ID         string `json:"id"`
	Task       string `json:"task"`
	IsDone     bool   `json:"is_done"`
	Created_at int64  `json:"created_at"`
}
