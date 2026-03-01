package handlers

import (
	"encoding/json"
	"net/http"
	"user-api/models"
	"user-api/repository"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод так себе, нужен POST", http.StatusMethodNotAllowed)
		return
	}

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Плохой запрос", http.StatusBadRequest)
		return
	}

	id, err := repository.CreateUser(u)
	if err != nil {
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})

}
