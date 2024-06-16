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
	_ = json.NewDecoder(r.Body).Decode(&user)

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
		err := json.NewEncoder(w).Encode(responses.Exception{Message: errLogin.Error()})
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if okLogin {

		claim := utils.JwtPayload{
			Email:  userService.User.Email,
			RoleID: userService.User.RoleID,
		}

		tokenString, expiration, err := utils.CreateToken(user.DocumentID, claim)
		if err != nil {
			utils.ErrorLogger.Println(err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(responses.Exception{Message: err.Error()})
			if err != nil {
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		userProfile := models.UserProfile{
			Email:     userService.User.Email,
			FirstName: userService.User.FirstName,
			LastName:  userService.User.LastName,
			Major:     userService.User.Major,
			RoleID:    userService.User.RoleID,
		}

		json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString, Expiration: expiration, UserProfile: userProfile})

	}
}
