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
	"time"

	"github.com/gorilla/mux"
)

var (
	matchService = services.MatchService{}
	scoreService = services.ScoreService{}
	UtilsService = services.UtilsService{}
)

func validateBodyParams(body []byte) (time.Time, int, int, int, *int, *int, error) {
    var requestBody struct {
        MatchDate      string `json:"match_date"`
        TeamLocalID    int    `json:"team_local_id"`
        TeamVisitorID  int    `json:"team_visitor_id"`
        ChampionshipID int    `json:"championship_id"`
        StageID        *int   `json:"stage_id"`
        GroupSID       *int   `json:"group_s_id"`
    }

    if err := json.Unmarshal(body, &requestBody); err != nil {
        return time.Time{}, 0, 0, 0, nil, nil, fmt.Errorf("error decoding request body: %v", err)
    }

    matchDate, errMatchDate := time.Parse(time.RFC3339, requestBody.MatchDate)
    if errMatchDate != nil {
        return time.Time{}, 0, 0, 0, nil, nil, fmt.Errorf("invalid match_date format: %v", errMatchDate)
    }

    return matchDate, requestBody.TeamLocalID, requestBody.TeamVisitorID, requestBody.ChampionshipID, requestBody.StageID, requestBody.GroupSID, nil
}

func validateEntities(championshipID, teamLocalID, teamVisitorID int) error {
	if ok, err := championshipService.ValidateChampionship(championshipID); err != nil || !ok {
		return fmt.Errorf("invalid championship_id %d: %v", championshipID, err)
	}
	if ok, err := teamService.CheckTeamExistsByID(teamLocalID); err != nil || !ok {
		return fmt.Errorf("invalid team_local_id %d: %v", teamLocalID, err)
	}
	if ok, err := teamService.CheckTeamExistsByID(teamVisitorID); err != nil || !ok {
		return fmt.Errorf("invalid team_visitor_id %d: %v", teamVisitorID, err)
	}
	return nil
}

func GetAllMatchesByChampionshipID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	championship_id, err := strconv.Atoi(vars["championship_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", err)
		return
	}
	if ok, err := championshipService.ValidateChampionship(championship_id); err != nil || !ok {
		utils.ErrorLogger.Printf("Error validating championship: %d: %v", championship_id, err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", err)
		return 
	}

	teams, err := matchService.GetAllMatchesByChampionshipID(championship_id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al obtener los equipos", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}

func GetAllPlayedMatchesByChampionshipID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	championship_id, err := strconv.Atoi(vars["championship_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", err)
		return
	}
	if ok, err := championshipService.ValidateChampionship(championship_id); err != nil || !ok {
		utils.ErrorLogger.Printf("Error validating championship: %d: %v", championship_id, err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", err)
		return 
	}

	teams, err := matchService.GetAllPlayedMatchesByChampionshipID(championship_id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al obtener los equipos", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}

func GetNotPlayedMatchesByChampionshipID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	championship_id, err := strconv.Atoi(vars["championship_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", err)
		return
	}
	if ok, err := championshipService.ValidateChampionship(championship_id); err != nil || !ok {
		utils.ErrorLogger.Printf("Error validating championship: %d: %v", championship_id, err)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", err)
		return 
	}

	teams, err := matchService.GetNotPlayedMatchesByChampionshipID(championship_id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al obtener los equipos", err)
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
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid match_id", err)
		return
	}

	result, err := matchService.GetMatchResult(matchID)
	if err != nil {
		if err.Error() == fmt.Sprintf("no match found for match_id: %d", matchID) {
			utils.RespondWithError(w, http.StatusNotFound, err.Error(), err)
		} else if err.Error() == fmt.Sprintf("match with match_id: %d has not been played yet", matchID) {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error(), err)
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error getting result", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func InsertMatch(w http.ResponseWriter, r *http.Request) {
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

	matchDate, teamLocalID, teamVisitorID, championshipID, stageID, groupSID, err := validateBodyParams(requestBody)
	//PRINT VARIABLES
	utils.InfoLogger.Println("MatchDate Controller: ", matchDate)
	utils.InfoLogger.Println("TeamLocalID Controller: ", teamLocalID)
	utils.InfoLogger.Println("TeamVisitorID Controller: ", teamVisitorID)
	utils.InfoLogger.Println("ChampionshipID Controller: ", championshipID)
	utils.InfoLogger.Println("StageID Controller: ", stageID)
	utils.InfoLogger.Println("GroupSID Controller: ", groupSID)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	if err := validateEntities(championshipID, teamLocalID, teamVisitorID); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	match := models.Match{
		MatchDate:      matchDate,
		TeamLocalID:    teamLocalID,
		TeamVisitorID:  teamVisitorID,
		ChampionshipID: championshipID,
		StageID: 	    stageID,
		GroupSID:       groupSID,
	}
	// PRINT MATCH
	utils.InfoLogger.Println("Match Controller: ", match)
	id, err := matchService.InsertMatch(match)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating match", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.Response{Data: "Match created successfully with id: " + strconv.Itoa(int(id))})
}

func UpdateMatch(w http.ResponseWriter, r *http.Request) {
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
		MatchID        int    `json:"match_id"`
		MatchDate      string `json:"match_date"`
		TeamLocalID    int    `json:"team_local_id"`
		TeamVisitorID  int    `json:"team_visitor_id"`
		GoalsLocal     int    `json:"goals_local"`
		GoalsVisitor   int    `json:"goals_visitor"`
		ChampionshipID int    `json:"championship_id"`
	}

	if err := json.Unmarshal(requestBody, &requestParams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	matchDate, errMatchDate := time.Parse(time.RFC3339, requestParams.MatchDate)
	if errMatchDate != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid match_date format", errMatchDate)
		return
	}

	matchID := requestParams.MatchID
	goalsLocal := requestParams.GoalsLocal
	goalsVisitor := requestParams.GoalsVisitor

	if err := validateEntities(requestParams.ChampionshipID, requestParams.TeamLocalID, requestParams.TeamVisitorID); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	match := models.Match{
		MatchID:        matchID,
		MatchDate:      matchDate,
		TeamLocalID:    requestParams.TeamLocalID,
		TeamVisitorID:  requestParams.TeamVisitorID,
		ChampionshipID: requestParams.ChampionshipID,
		GoalsLocal:     &goalsLocal,
		GoalsVisitor:   &goalsVisitor,
	}

	id, err := matchService.UpdateMatch(match)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error updating match", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Response{Data: "Match updated successfully with id: " + strconv.Itoa(int(id))})
}

func DeleteMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["match_id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid match_id", err)
		return
	}

	id, err := matchService.DeleteMatch(matchID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error deleting match", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Response{Data: "Match deleted successfully with id: " + strconv.Itoa(int(id))})
}

func InsertResult(w http.ResponseWriter, r *http.Request) {
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
		MatchID      *int `json:"match_id"`
		GoalsLocal   *int `json:"goals_local"`
		GoalsVisitor *int `json:"goals_visitor"`
	}

	if err := json.Unmarshal(requestBody, &requestParams); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	matchID := *requestParams.MatchID
	goalsLocal := *requestParams.GoalsLocal
	goalsVisitor := *requestParams.GoalsVisitor

	utils.InfoLogger.Println("MatchID Controller: ", matchID)
	utils.InfoLogger.Println("GoalsLocal Controller: ", goalsLocal)
	utils.InfoLogger.Println("GoalsVisitor Controller: ", goalsVisitor)

	if requestParams.MatchID == nil|| requestParams.GoalsLocal == nil || requestParams.GoalsVisitor == nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body", fmt.Errorf("%v %v %v", matchID, goalsLocal, goalsVisitor))
		return
	}

	matchData := models.Match{
		MatchID:      matchID,
		GoalsLocal:   &goalsLocal,
		GoalsVisitor: &goalsVisitor,
	}

	match, err := matchService.ValidateMatch(matchID)
	if err != nil || !match {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid match_id, this match not exist", err)
		return
	}

	id, err := matchService.InsertResult(matchData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error inserting result", err)
		return
	} else {
		utils.InfoLogger.Println("Result inserted successfully to match_id: ", id)
	}

	err = calculateAndAssignPoints(matchID, goalsLocal, goalsVisitor)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error calculating and assigning points", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.Response{Data: "Result inserted successfully with id: " + strconv.Itoa(int(id))})
}

func calculateAndAssignPoints(matchID, goalsLocal, goalsVisitor int) error {
	predictions, err := predictionService.GetPredictionsByMatchID(matchID)
	if err != nil {
		return fmt.Errorf("error getting predictions: %v", err)
	}

	for _, prediction := range predictions {
		points, err := calculatePoints(prediction.GoalsLocal, prediction.GoalsVisitor, goalsLocal, goalsVisitor)
		if err != nil {
			return fmt.Errorf("error calculating points for user %v: %v", prediction.DocumentID, err)
		}
		err = scoreService.InsertOrUpdateScore(prediction.DocumentID, matchID, points)
		if err != nil {
			return fmt.Errorf("error updating score for user %v: %v", prediction.DocumentID, err)
		}
	}

	return nil
}

func calculatePoints(predictedLocal, predictedVisitor, actualLocal, actualVisitor int) (int, error) {
	if predictedLocal == actualLocal && predictedVisitor == actualVisitor {
		return UtilsService.GetPointsExactResult() 
	} else if (predictedLocal > predictedVisitor && actualLocal > actualVisitor) ||
		(predictedLocal < predictedVisitor && actualLocal < actualVisitor) ||
		(predictedLocal == predictedVisitor && actualLocal == actualVisitor) {
		return UtilsService.GetPointsCorrectResult() 
	} else {
		return 0, nil 
	}
}
