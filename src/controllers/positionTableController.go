package controllers

import (
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	positionTableService = &services.PositionTableService{}
)

func GetPositionTableByChampionship(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	championshipID, err := strconv.Atoi(vars["championship_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", err)
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

func GetUserScoresByGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user_id", err)
		return
	}
	groupID, err := strconv.Atoi(vars["group_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid group_id", err)
		return
	}

	userScores, err := positionTableService.GetUserScoresByGroup(userID, groupID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error retrieving user scores", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userScores)
}
