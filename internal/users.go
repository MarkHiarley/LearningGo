package users

import (
	"API/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	data := []models.User{
		{
			Id:   1,
			Name: "teste,",
			Age:  12,
		},
		{
			Id:   1,
			Name: "teste,",
			Age:  12,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Printf("Erro ao encodar JSON: %v", err)
	}
}
