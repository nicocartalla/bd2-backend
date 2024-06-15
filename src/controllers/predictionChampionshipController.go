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
)

var (
	predictionChampionshipService = &services.PredictionChampionshipService{}
)


func GetPredictionChampionshipByUserAndChampionshipID(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		DocumentID    string `json:"document_id"`
		ChampionshipID int   `json:"championship_id"`
	}

	if r.Body != nil {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}
		if err := json.Unmarshal(body, &requestBody); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
			return
		}
	}

	predictions, err := predictionService.GetPredictionsByUserAndChampionshipID(requestBody.DocumentID, requestBody.ChampionshipID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error retrieving predictions", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(predictions)
}



func InsertPredictionChampionship(w http.ResponseWriter, r *http.Request) {
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
		Champion      *int   `json:"champion"`
		SubChampion   *int   `json:"subchampion"`
		ChampionshipID *int   `json:"championship_id"`
	}

	if err := json.Unmarshal(requestBody, &requestParams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	if requestParams.DocumentID == "" || requestParams.Champion == nil || requestParams.SubChampion == nil || requestParams.ChampionshipID == nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request parameters", fmt.Errorf("invalid parameters: %v", requestParams))
		return
	}

	predictionChampionship := models.PredictionChampionship{
		DocumentID:     requestParams.DocumentID,
		Champion:       *requestParams.Champion,
		SubChampion:    *requestParams.SubChampion,
		ChampionshipID: *requestParams.ChampionshipID,
	}

	ok := userService.CheckUserExistsByDocumentID(predictionChampionship.DocumentID)
	if !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "User does not exist", fmt.Errorf("%s", predictionChampionship.DocumentID))
		return
	}

	if ok, err := championshipService.ValidateChampionship(predictionChampionship.ChampionshipID); err != nil || !ok {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", fmt.Errorf("%d: %v", predictionChampionship.ChampionshipID, err))
		return
	}

	id, err := predictionChampionshipService.InsertPredictionChampionship(predictionChampionship)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error inserting prediction Championship", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.Response{Data: "Prediction Championship created successfully with id: " + strconv.Itoa(int(id))})
	utils.InfoLogger.Println("Prediction Championship created successfully with id: ", id)
}
	