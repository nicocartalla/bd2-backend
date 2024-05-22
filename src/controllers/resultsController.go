package controllers

import (
	"bd2-backend/src/utils"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"encoding/json"
	"net/http"
)



var (
	resultService = &services.ResultService{}
)

func GetAllResults(w http.ResponseWriter, r *http.Request) {
	teams, err := resultService.CheckAllResults()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		utils.ErrorLogger.Println(err)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Error al obtener los equipos"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}
