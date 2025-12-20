package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"my-todo/database"
	"my-todo/handlers"
	"my-todo/middleware"
)

func main() {
	database.Connect()
	database.Migrate()

	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	todoRouter := router.PathPrefix("/").Subrouter()
	todoRouter.Use(middleware.Auth)

	todoRouter.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	todoRouter.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	todoRouter.HandleFunc("/todo", handlers.GetTodoByID).Methods("GET")
	todoRouter.HandleFunc("/todo", handlers.DeleteTodo).Methods("DELETE")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
