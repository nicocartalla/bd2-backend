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
    "io"
    "github.com/gorilla/mux"
)

var (
    championshipService = services.ChampionshipService{}
    ScoreChampionshipService = services.ScoreChampionshipService{}
)

func GetAllChampionships(w http.ResponseWriter, r *http.Request) {
    championships, err := championshipService.GetAllChampionships()
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error getting championships", err)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(championships)
}

func GetChampionshipByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship ID", err)
        return
    }

    championship, err := championshipService.GetChampionshipByID(id)
    if err != nil {
        utils.RespondWithError(w, http.StatusNotFound, err.Error(), err)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(championship)
}

func CreateChampionship(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Name             string `json:"name"`
        Year             int    `json:"year"`
        Country          string `json:"country"`
        ChampionshipType string `json:"championship_type"`
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

    championship := models.Championship{
        Name:             requestBody.Name,
        Year:             requestBody.Year,
        Country:          requestBody.Country,
        ChampionshipType: requestBody.ChampionshipType,
    }

    id, err := championshipService.CreateChampionship(championship)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error creating championship", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(responses.Response{Data: fmt.Sprintf("Championship created successfully with id: %d", id)})
}

func UpdateChampionship(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship ID", err)
        return
    }

    var championship models.Championship
    err = json.NewDecoder(r.Body).Decode(&championship)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    _, err = championshipService.UpdateChampionship(id, championship)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error updating championship", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(responses.Response{Data: "Championship updated successfully"})
}

func DeleteChampionship(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship ID", err)
        return
    }

    _, err = championshipService.DeleteChampionship(id)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error deleting championship", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(responses.Response{Data: "Championship deleted successfully"})
}
// funcion para setear el championship y los champions, recibe en el body el champio y el sub_champion y el championhip_id
func SetChampionshipChampions(w http.ResponseWriter, r *http.Request) {
    var requestBody struct {
        ChampionID   int `json:"champion_id"`
        SubChampionID int `json:"sub_champion_id"`
        ChampionshipID int `json:"championship_id"`
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

    if ok, err := championshipService.ValidateChampionship(requestBody.ChampionshipID); err != nil || !ok {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid championship_id", fmt.Errorf("%d: %v", requestBody.ChampionshipID, err))
        return
    }

    if requestBody.ChampionID == requestBody.SubChampionID {
        utils.RespondWithError(w, http.StatusBadRequest, "Champion and subchampion must be different", fmt.Errorf("champion_id: %d, sub_champion_id: %d", requestBody.ChampionID, requestBody.SubChampionID))
        return
    }

    _, err := championshipService.SetChampionshipChampions(requestBody.ChampionshipID, requestBody.ChampionID, requestBody.SubChampionID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error updating championship", err)
        return
    } else {
        utils.InfoLogger.Println("Championship champions updated successfully")
    }

    err = CalculateAndAssignChampionshipPoints(requestBody.ChampionshipID, requestBody.ChampionID, requestBody.SubChampionID)
    if err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, "Error calculating points", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(responses.Response{Data: "Championship champions updated successfully"})
}

//calcular y asignar los puntos a los usuarios que acertaron campeon y subcampeon
func CalculateAndAssignChampionshipPoints(championshipID, championID, subChampionID int) error {
    predictions, err := predictionChampionshipService.GetPredictionsChampionshipByChampionshipID(championshipID)
    if err != nil {
        return fmt.Errorf("error getting predictions: %v", err)
    }

    for _, prediction := range predictions {
        points, err := calculateChampionshipPoints(prediction.Champion, prediction.SubChampion, championID, subChampionID)
        if err != nil {
            return fmt.Errorf("error calculating points for user %v: %v", prediction.DocumentID, err)
        }
        err = ScoreChampionshipService.InsertOrUpdateScoreChampionship(prediction.DocumentID, championshipID, points)
        if err != nil {
            return fmt.Errorf("error updating score for user %v: %v", prediction.DocumentID, err)
        }
    }

    return nil
}

func calculateChampionshipPoints(champion, subChampion, championID, subChampionID int) (int, error) {
    if champion == championID && subChampion == subChampionID {
        return UtilsService.GetPointsChampion()
    } else if (champion == subChampionID && subChampion == championID) ||
        (champion == championID && subChampion == subChampionID) {
        return UtilsService.GetPointsSubChampion()
    } else {
        return 0, nil
    }
}
