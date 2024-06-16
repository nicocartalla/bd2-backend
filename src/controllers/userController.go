package controllers

import (
	"bd2-backend/src/models"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"fmt"
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

// Add user to championship
func AddUserToChampionship(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		DocumentID     string `json:"document_id"`
		ChampionshipID int    `json:"championship_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorLogger.Println(err)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al decodificar la solicitud"})
		return
	}

	//Validate ChampionshipID
	if ok, err := championshipService.ValidateChampionship(requestBody.ChampionshipID); err != nil || !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", fmt.Errorf("%d: %v", requestBody.ChampionshipID, err))
		return
	}

	var userService = &services.UserService{}
	errCreate := userService.AddUserToChampionship(requestBody.DocumentID, requestBody.ChampionshipID)
	if errCreate != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		utils.ErrorLogger.Println(errCreate)
		json.NewEncoder(w).Encode(responses.Exception{Message: errCreate.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"document_id":     requestBody.DocumentID,
		"championship_id": requestBody.ChampionshipID,
		"status":          "User added to championship",
	})
}
