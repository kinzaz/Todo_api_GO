package task

type TaskCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
