package controllers

import (
	"bd2-backend/src/utils"
	"net/http"
)

// validate admin user
func ValidateAdminUser(w http.ResponseWriter, r *http.Request) (bool) {

	payload, err := utils.GetJwtPayloadFromClaim(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error getting claims from context", err)
	}

	if payload.RoleID == 1 {
		return true
	}

	utils.RespondWithError(w, http.StatusInternalServerError, "Unauthorized", err)
	return false

}
