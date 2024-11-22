package task

import (
	"fmt"
	"net/http"
	"strconv"
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
	router.HandleFunc("GET /task/{id}", handler.GetTask())
	router.HandleFunc("PATCH /task/{id}", handler.Update())

	router.HandleFunc("GET /tasks", handler.GetAll())
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
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
		}

		_, err = handler.TaskRepository.GetById(uint(id))

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

func (handler *TaskHandler) GetTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
		}
		task, err := handler.TaskRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, "Задачи не существует", http.StatusBadRequest)
			return
		}
		response.Json(w, task, 200)
	}
}

func (handler *TaskHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
		}
		if idString == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}

		body, err := request.HandleBody[TaskUpdateRequest](w, r)
		if err != nil {
			return
		}

		existedTask, err := handler.TaskRepository.GetById(uint(id))
		fmt.Println(existedTask)
		if err != nil {
			http.Error(w, "task not found", http.StatusNotFound)
			return
		}
		existedTask.Title = body.Title
		existedTask.Description = body.Description
		existedTask.Completed = body.Completed

		updatedTask, err := handler.TaskRepository.Update(existedTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, updatedTask, 200)
	}
}

func (handler *TaskHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := handler.TaskRepository.GetAll()
		if err != nil {
			http.Error(w, "failed to fetch tasks: "+err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, tasks, 200)
	}
}
