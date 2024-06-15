package controllers

import (
	"encoding/json"
	"net/http"
	"bd2-backend/src/utils"
)


func Ping(w http.ResponseWriter, r *http.Request) {

	payload, err := utils.GetJwtPayloadFromClaim(r.Context())
	if err != nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("error getting claims from context")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"user": payload.Email,
		"role_id": payload.RoleID,
		"message": "pong",
	}

	json.NewEncoder(w).Encode(response)
}
