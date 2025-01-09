package internal

import "groot/internal/models/tasks"

type ListResponse struct {
	Code    int          `json:"code"`
	Data    []tasks.Task `json:"data"`
	Error   string       `json:"error"`
	Message string       `json:"message"`
}
