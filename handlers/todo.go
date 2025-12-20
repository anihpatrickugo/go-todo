package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"my-todo/database"
	"my-todo/models"
    "my-todo/middleware"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(uint)

	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.UserID = userID
	database.DB.Create(&todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(uint)

	var todos []models.Todo
	database.DB.Where("user_id = ?", userID).Find(&todos)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(uint)
	
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var todo models.Todo
	result := database.DB.First(&todo, id, "user_id = ?", userID)
	if result.Error != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(uint)

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	database.DB.Delete(&models.Todo{}, id, "user_id = ?", userID)
	w.WriteHeader(http.StatusNoContent)
}
