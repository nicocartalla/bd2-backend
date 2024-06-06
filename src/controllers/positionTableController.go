package controllers

import (
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
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

