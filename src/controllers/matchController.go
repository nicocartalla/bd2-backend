package controllers

import (
	"bd2-backend/src/models"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var (
	matchService = services.MatchService{}
)

func respondWithError(w http.ResponseWriter, code int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	utils.ErrorLogger.Println(message, err)
	json.NewEncoder(w).Encode(responses.Exception{Message: message})
}

func validateQueryParams(queryParams map[string]string) (time.Time, int, int, int, error) {
	matchDate, errMatchDate := time.Parse(time.RFC3339, queryParams["match_date"])
	teamLocalID, errTeamLocalID := strconv.Atoi(queryParams["team_local_id"])
	teamVisitorID, errVisitorID := strconv.Atoi(queryParams["team_visitor_id"])
	championshipID, errChampionshipID := strconv.Atoi(queryParams["championship_id"])

	if errMatchDate != nil || errTeamLocalID != nil || errVisitorID != nil || errChampionshipID != nil {
		return time.Time{}, 0, 0, 0, fmt.Errorf("invalid query parameters: %v %v %v %v", errMatchDate, errTeamLocalID, errVisitorID, errChampionshipID)
	}
	return matchDate, teamLocalID, teamVisitorID, championshipID, nil
}

func validateEntities(championshipID, teamLocalID, teamVisitorID int) error {
	if ok, err := matchService.ValidateChampionship(championshipID); err != nil || !ok {
		return fmt.Errorf("invalid championship_id: %v", err)
	}
	if ok, err := teamService.CheckTeamExistsByID(teamLocalID); err != nil || !ok {
		return fmt.Errorf("invalid team_local_id: %v", err)
	}
	if ok, err := teamService.CheckTeamExistsByID(teamVisitorID); err != nil || !ok {
		return fmt.Errorf("invalid team_visitor_id: %v", err)
	}
	return nil
}

func GetAllMatchResults(w http.ResponseWriter, r *http.Request) {
	teams, err := matchService.GetAllMatchResults()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error al obtener los equipos", err)
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
		respondWithError(w, http.StatusBadRequest, "Invalid match_id", err)
		return
	}

	result, err := matchService.GetMatchResult(matchID)
	if err != nil {
		if err.Error() == fmt.Sprintf("no match found for match_id: %d", matchID) {
			respondWithError(w, http.StatusNotFound, err.Error(), err)
		} else if err.Error() == fmt.Sprintf("match with match_id: %d has not been played yet", matchID) {
			respondWithError(w, http.StatusBadRequest, err.Error(), err)
		} else {
			respondWithError(w, http.StatusInternalServerError, "Error getting result", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func InsertMatch(w http.ResponseWriter, r *http.Request) {
	queryParams := map[string]string{
		"match_date":     r.URL.Query().Get("match_date"),
		"team_local_id":  r.URL.Query().Get("team_local_id"),
		"team_visitor_id": r.URL.Query().Get("team_visitor_id"),
		"championship_id": r.URL.Query().Get("championship_id"),
	}

	matchDate, teamLocalID, teamVisitorID, championshipID, err := validateQueryParams(queryParams)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid query parameters", err)
		return
	}

	if err := validateEntities(championshipID, teamLocalID, teamVisitorID); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	match := models.Match{
		MatchDate:      matchDate,
		TeamLocalID:    teamLocalID,
		TeamVisitorID:  teamVisitorID,
		ChampionshipID: championshipID,
	}

	id, err := matchService.InsertMatch(match)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating match", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.Response{Data: "Match created successfully with id: " + strconv.Itoa(int(id))})
}

func UpdateMatch(w http.ResponseWriter, r *http.Request) {
	queryParams := map[string]string{
		"match_id":       r.URL.Query().Get("match_id"),
		"match_date":     r.URL.Query().Get("match_date"),
		"team_local_id":  r.URL.Query().Get("team_local_id"),
		"team_visitor_id": r.URL.Query().Get("team_visitor_id"),
		"goals_local":    r.URL.Query().Get("goals_local"),
		"goals_visitor":  r.URL.Query().Get("goals_visitor"),
		"championship_id": r.URL.Query().Get("championship_id"),
	}

	matchID, errMatchID := strconv.Atoi(queryParams["match_id"])
	goalsLocal, errGoalsLocal := strconv.Atoi(queryParams["goals_local"])
	goalsVisitor, errGoalsVisitor := strconv.Atoi(queryParams["goals_visitor"])

	matchDate, teamLocalID, teamVisitorID, championshipID, err := validateQueryParams(queryParams)
	if err != nil || errMatchID != nil || errGoalsLocal != nil || errGoalsVisitor != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid query parameters", fmt.Errorf("%v %v %v %v", err, errMatchID, errGoalsLocal, errGoalsVisitor))
		return
	}

	if err := validateEntities(championshipID, teamLocalID, teamVisitorID); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	match := models.Match{
		MatchID:        matchID,
		MatchDate:      matchDate,
		TeamLocalID:    teamLocalID,
		TeamVisitorID:  teamVisitorID,
		ChampionshipID: championshipID,
		GoalsLocal:     &goalsLocal,
		GoalsVisitor:   &goalsVisitor,
	}

	id, err := matchService.UpdateMatch(match)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating match", err)
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
		respondWithError(w, http.StatusBadRequest, "Invalid match_id", err)
		return
	}

	id, err := matchService.DeleteMatch(matchID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error deleting match", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Response{Data: "Match deleted successfully with id: " + strconv.Itoa(int(id))})
}

func InsertResult(w http.ResponseWriter, r *http.Request) {
	queryParams := map[string]string{
		"match_id":      r.URL.Query().Get("match_id"),
		"goals_local":   r.URL.Query().Get("goals_local"),
		"goals_visitor": r.URL.Query().Get("goals_visitor"),
	}

	matchID, errMatchID := strconv.Atoi(queryParams["match_id"])
	goalsLocal, errGoalsLocal := strconv.Atoi(queryParams["goals_local"])
	goalsVisitor, errGoalsVisitor := strconv.Atoi(queryParams["goals_visitor"])

	if errMatchID != nil || errGoalsLocal != nil || errGoalsVisitor != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid query parameters", fmt.Errorf("%v %v %v", errMatchID, errGoalsLocal, errGoalsVisitor))
		return
	}
	matchData := models.Match{
		MatchID:      matchID,
		GoalsLocal:   &goalsLocal,
		GoalsVisitor: &goalsVisitor,
	}

	match, err := matchService.ValidateMatch(matchID)
	if err != nil || !match {
		respondWithError(w, http.StatusBadRequest, "Invalid match_id, this match not exist", err)
		return
	}

	id, err := matchService.InsertResult(matchData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error inserting result", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responses.Response{Data: "Result inserted successfully with id: " + strconv.Itoa(int(id))})
}

func GetMatchesNotPlayedYet(w http.ResponseWriter, r *http.Request) {
	matches, err := matchService.GetMatchesNotPlayedYet()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting matches", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}

