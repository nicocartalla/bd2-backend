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

    "github.com/gorilla/mux"
)

var (
    championshipService = services.ChampionshipService{}
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
    var championship models.Championship
    err := json.NewDecoder(r.Body).Decode(&championship)
    if err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
        return
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
