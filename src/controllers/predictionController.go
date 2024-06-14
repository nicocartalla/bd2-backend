package controllers

import (
	"bd2-backend/src/models"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	predictionService = &services.PredictionService{}
)

func GetPredictionsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["document_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid document_id", err)
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
	var requestBody []byte
	if r.Body != nil {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}
		requestBody = body
	}

	var requestParams struct {
		DocumentID   string `json:"document_id"`
		MatchID      *int   `json:"match_id"`
		GoalsLocal   *int   `json:"goals_local"`
		GoalsVisitor *int   `json:"goals_visitor"`
		GroupID      *int   `json:"group_id"`
	}

	if err := json.Unmarshal(requestBody, &requestParams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	if requestParams.DocumentID == "" || requestParams.MatchID == nil || requestParams.GoalsLocal == nil || requestParams.GoalsVisitor == nil || requestParams.GroupID == nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request parameters", fmt.Errorf("invalid parameters: %v", requestParams))
		return
	}

	prediction := models.Prediction{
		DocumentID:   requestParams.DocumentID,
		MatchID:      *requestParams.MatchID,
		GoalsLocal:   *requestParams.GoalsLocal,
		GoalsVisitor: *requestParams.GoalsVisitor,
		GroupID:      *requestParams.GroupID,
	}

	if ok, err := matchService.ValidateMatch(prediction.MatchID); err != nil || !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid match_id", fmt.Errorf("%d: %v", prediction.MatchID, err))
		return
	}

	if ok, err := matchService.IsMatchUpcoming(prediction.MatchID); err != nil || !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Match has been played or the hours until played is less than permitted", fmt.Errorf("%d: %v", prediction.MatchID, err))
		return
	}

	id, err := predictionService.InsertPrediction(prediction)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error inserting prediction", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.Response{Data: "Prediction created successfully with id: " + strconv.Itoa(int(id))})
	utils.InfoLogger.Println("Prediction created successfully with id: ", id)
}
