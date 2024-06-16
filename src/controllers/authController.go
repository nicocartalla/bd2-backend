package controllers

import (
	"bd2-backend/src/models"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userService services.UserService
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: "Invalid request payload"})
		return
	}

	if user.Email == "" || user.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responses.Exception{Message: "email and password are required"})
		return
	}

	userService.User = user
	okLogin, errLogin := userService.ValidateLogin()
	if errLogin != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		utils.ErrorLogger.Println(errLogin.Error())
		json.NewEncoder(w).Encode(responses.Exception{Message: errLogin.Error()})
		return
	}

	if okLogin {
		claim := utils.JwtPayload{
			Email:  userService.User.Email,
			RoleID: userService.User.RoleID,
		}

		tokenString, expiration, err := utils.CreateToken(userService.User.DocumentID, claim)
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
			return
		}

		userProfile := models.UserProfile{
			DocumentID: userService.User.DocumentID,
			Email:      userService.User.Email,
			FirstName:  userService.User.FirstName,
			LastName:   userService.User.LastName,
			Major:      userService.User.Major,
			RoleID:     userService.User.RoleID,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.JwtToken{
			Token:       tokenString,
			Expiration:  expiration,
			UserProfile: userProfile,
		})
	}
}
