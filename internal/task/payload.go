package task

type TaskCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
