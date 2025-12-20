package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"my-todo/database"
	"my-todo/models"
	"my-todo/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User

	json.NewDecoder(r.Body).Decode(&input)

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateToken(user.ID)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
