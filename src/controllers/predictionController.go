package controllers

import (
	"bd2-backend/src/models"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
)

var (
	predictionService = &services.PredictionService{}
)

func GetPredictionsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid user_id", err)
		return
	}

	predictions, err := predictionService.GetPredictionsByUser(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error retrieving predictions", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(predictions)
}

func InsertPrediction(w http.ResponseWriter, r *http.Request) {
	queryParams := map[string]string{
		"document_id":      r.URL.Query().Get("document_id"),
		"match_id":     r.URL.Query().Get("match_id"),
		"goals_local":  r.URL.Query().Get("goals_local"),
		"goals_visitor": r.URL.Query().Get("goals_visitor"),
		"group_id":     r.URL.Query().Get("group_id"),
	}

	documentID := queryParams["document_id"]
	matchID, errMatchID := strconv.Atoi(queryParams["match_id"])
	goalsLocal, errGoalsLocal := strconv.Atoi(queryParams["goals_local"])
	goalsVisitor, errGoalsVisitor := strconv.Atoi(queryParams["goals_visitor"])
	groupID, errGroupID := strconv.Atoi(queryParams["group_id"])

	if errMatchID != nil || errGoalsLocal != nil || errGoalsVisitor != nil || errGroupID != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid query parameters", fmt.Errorf("%v %v %v %v", errMatchID, errGoalsLocal, errGoalsVisitor, errGroupID))
		return
	}

	prediction := models.Prediction{
		DocumentID:   documentID,
		MatchID:      matchID,
		GoalsLocal:   goalsLocal,
		GoalsVisitor: goalsVisitor,
		GroupID:      groupID,
	}

	id, err := predictionService.InsertPrediction(prediction)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error inserting prediction", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.Response{Data: "Prediction created successfully with id: " + strconv.Itoa(int(id))})
}
