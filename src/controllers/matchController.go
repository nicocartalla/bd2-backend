package controllers

import (
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
 	"fmt"
	"bd2-backend/src/models"
	"time"
)

var (
	matchService = &services.MatchService{}
)

func GetAllMatchResults(w http.ResponseWriter, r *http.Request) {
	teams, err := matchService.CheckAllResults()
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

func GetMatchResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["match_id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorLogger.Println("Invalid match_id:", err)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid match_id"})
		return
	}

	result, err := matchService.GetMatchResult(matchID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if err.Error() == fmt.Sprintf("no match found for match_id: %d", matchID) {
			w.WriteHeader(http.StatusNotFound)
		} else if err.Error() == fmt.Sprintf("match with match_id: %d has not been played yet", matchID) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		utils.ErrorLogger.Println("Error getting result:", err)
		json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func InsertMatch(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	matchDateStr := queryParams.Get("match_date")
	teamLocalIDStr := queryParams.Get("team_local_id")
	teamVisitorIDStr := queryParams.Get("team_visitor_id")
	championshipIDStr := queryParams.Get("championship_id")

		matchDate, errMatchDate := time.Parse(time.RFC3339, matchDateStr)
		teamLocalID, errTeamLocalID := strconv.Atoi(teamLocalIDStr)
		teamVisitorID, errVisitorID := strconv.Atoi(teamVisitorIDStr)
		championshipID, errChampionshioID := strconv.Atoi(championshipIDStr)

		if errMatchDate != nil || errTeamLocalID != nil || errVisitorID != nil || errChampionshioID != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ErrorLogger.Println("Invalid query parameters:", errMatchDate, errTeamLocalID, errVisitorID, errChampionshioID)
			json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid query parameters"})
			return
		}

		okChampionship, err := matchService.ValidateChampionship(championshipID)
		if err != nil || !okChampionship {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ErrorLogger.Println("Invalid championship_id:", err)
			json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid championship_id"})
			return
		}

		okTeamLocal, err := teamService.CheckTeamExistsByID(teamLocalID)
		if err != nil || !okTeamLocal {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ErrorLogger.Println("Invalid team_local_id:", err)
			json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid team_local_id"})
			return
		}

		okTeamVisitor, err := teamService.CheckTeamExistsByID(teamVisitorID)
		if err != nil || !okTeamVisitor {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			utils.ErrorLogger.Println("Invalid team_visitor_id:", err)
			json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid team_visitor_id"})
			return
		}

		match := models.Match{
			MatchDate:      matchDate,
			TeamLocalID:    teamLocalID,
			TeamVisitorID:  teamVisitorID,
			ChampionshipID: championshipID,
		}
	
		id, err := matchService.InsertMatch(match)
		if err != nil || id != 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			utils.ErrorLogger.Println("Error creating match:", err)
			json.NewEncoder(w).Encode(responses.Exception{Message: "Error creating match"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responses.Response{Data: "Match created successfully"})
	
	}