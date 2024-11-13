package main

import (
	"fmt"
	"log"
	"net/http"
	"todo_GO/configs"
)

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":" + conf.Port,
		Handler: router,
	}

	fmt.Printf("Server is running on %s port", conf.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
