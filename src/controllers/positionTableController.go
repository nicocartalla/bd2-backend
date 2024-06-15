package controllers

import (
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)

var (
	positionTableService = &services.PositionTableService{}
)

/* func GetPositionTableByChampionship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	championshipID, err := strconv.Atoi(vars["championship_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", err)
		return
	}
	//validate championship_id
	if ok, err := championshipService.ValidateChampionship(championshipID); err != nil || !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", fmt.Errorf("%d: %v", championshipID, err))
		return
	}
	positionTable, err := positionTableService.GetPositionTableByChampionship(championshipID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error retrieving leaderboard", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(positionTable)
}

 */

 func GetPositionTableByChampionship(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		ChampionshipID int `json:"championship_id"`
	}
	// Read and parse the request body
	if r.Body != nil {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}
		
		// Log the raw body for debugging
		utils.InfoLogger.Println("Raw request body: ", string(body))
		
		// Unmarshal the JSON payload
		if err := json.Unmarshal(body, &requestBody); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
			return
		}
	} else {
		utils.RespondWithError(w, http.StatusBadRequest, "Request body is empty", nil)
		return
	}
	

	championshipID := requestBody.ChampionshipID
	if championshipID == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "championship_id is required", nil)
		return
	}

	// Validate championship_id
	if ok, err := championshipService.ValidateChampionship(championshipID); err != nil || !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", fmt.Errorf("%d: %v", championshipID, err))
		return
	}

	positionTable, err := positionTableService.GetPositionTableByChampionship(championshipID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error retrieving leaderboard", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(positionTable)
}
