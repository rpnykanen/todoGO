package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	todo := Todo{NewStorage()}

	mux.HandleFunc("GET /todo", todo.Todo)
	mux.HandleFunc("POST /todo", todo.Create)
	mux.HandleFunc("GET /todo/{id}", todo.Read)

	mux.HandleFunc("PATCH /todo/{id}", todo.Update)
	mux.HandleFunc("DELETE /todo/{id}", todo.Delete)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("asd", err)
		log.Fatal(err)
	}
}
