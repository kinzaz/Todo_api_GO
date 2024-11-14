package main

import (
	"fmt"
	"log"
	"net/http"
	"todo_GO/configs"
	"todo_GO/internal/task"
	"todo_GO/pkg/db"
)

func main() {
	database := db.NewDb()
	conf := configs.LoadConfig()

	router := http.NewServeMux()

	/* Repositories */
	taskRepository := task.NewTaskRepository(database)

	/* Handlers */
	task.NewTaskHandler(router, task.TaskHandlerDeps{
		TaskRepository: taskRepository,
	})

	server := http.Server{
		Addr:    ":" + conf.Port,
		Handler: router,
	}

	fmt.Printf("Server is running on %s port", conf.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
