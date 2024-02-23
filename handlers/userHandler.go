package handlers

import (
	"encoding/json"
	"fmt"
	"grpc-go2/models"
	"io"
	"log"
	"net/http"
)

func (h handler) AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	var user models.User
	json.Unmarshal(body, &user)
	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
