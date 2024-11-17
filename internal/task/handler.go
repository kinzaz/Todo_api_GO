package task

import (
	"net/http"
	"todo_GO/pkg/request"
	"todo_GO/pkg/response"
)

type TaskHandler struct {
	TaskRepository *TaskRepository
}

type TaskHandlerDeps struct {
	TaskRepository *TaskRepository
}

func NewTaskHandler(router *http.ServeMux, deps TaskHandlerDeps) {
	handler := &TaskHandler{
		TaskRepository: deps.TaskRepository,
	}

	router.HandleFunc("POST /task", handler.Create())
	router.HandleFunc("DELETE /task/{id}", handler.Delete())
}

func (handler *TaskHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[TaskCreateRequest](w, r)

		task := NewTask(body.Title, body.Description)

		if err != nil {
			return
		}

		createdTask, err := handler.TaskRepository.Create(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdTask, 201)
	}
}

func (handler *TaskHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")

		_, err := handler.TaskRepository.GetById(idString)

		if err != nil {
			http.Error(w, "Задачи не существует", http.StatusBadRequest)
			return
		}

		err = handler.TaskRepository.Delete(idString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, nil, 200)
	}
}
