package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "my-todo/database"
	"my-todo/handlers"
)



func main() {
    database.Connect()
	database.Migrate()

    router := mux.NewRouter()

    router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
    router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
    router.HandleFunc("/todo", handlers.GetTodoByID).Methods("GET")
    router.HandleFunc("/todo", handlers.DeleteTodo).Methods("DELETE")

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
