package controllers

import (
	"bd2-backend/src/models"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"net/http"
)

type CreateUserResponse models.CreateUserResponse


func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	var userService = &services.UserService{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorLogger.Println(err)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al decodificar el usuario"})
		return
	}

//	id, errCreate := user.CreateUser()
	userService.User = user
	docId, errCreate := userService.CreateUser()
	if errCreate != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		utils.ErrorLogger.Println(errCreate)
		json.NewEncoder(w).Encode(responses.Exception{Message: errCreate.Error()})
		return
	}

	if docId == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorLogger.Println(errCreate)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al crear el usuario"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{DocumentID: docId, Status: "User created"})
}
