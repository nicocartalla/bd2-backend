package controllers

import (
	"bd2-backend/src/utils"
	"bd2-backend/src/services"
	"net/http"
)

// validate admin user
func ValidateAdminUser(w http.ResponseWriter, r *http.Request) (bool) {

	payload, err := utils.GetJwtPayloadFromClaim(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error getting claims from context", err)
	}

	roleService := services.RoleService{}
	roleData, err := roleService.GetRoleId("Admin")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error getting role id", err)
		return false
	}
	if payload.RoleID == roleData.ID {
		return true
	}

	utils.RespondWithError(w, http.StatusInternalServerError, "Unauthorized", err)
	return false

}
