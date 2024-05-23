package controllers

import (
	"bd2-backend/src/utils"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type createTeamResponse struct {
	ID		int    `json:"id"`
	Name	string `json:"name"`
}

var (
	teamService = &services.TeamService{}
)

func GetTeams(w http.ResponseWriter, r *http.Request) {
	var teamService = &services.TeamService{}
	teams, err := teamService.GetTeams()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		utils.ErrorLogger.Println(err)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener los equipos"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}

func CheckTeamExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamName := vars["name"]
	exists := teamService.CheckTeamExistsByName(teamName)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"exists": exists})
}