package controllers

import (
	"bd2-backend/src/config"
	"bd2-backend/src/models"
	"bd2-backend/src/responses"
	"bd2-backend/src/services"
	"bd2-backend/src/utils"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

//var jwtToken = []byte("secret")
var jwtToken []byte

//obtengo la clave para generar el token
func init() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		utils.ErrorLogger.Fatal("cannot load config:", err)
	}
	jwtToken = []byte(cfg.JwtKey)

}

// CreateToken crea un token JWT
func CreateToken(w http.ResponseWriter, r *http.Request) {

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
		expiration := time.Now().Add(time.Hour * time.Duration(1)).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": 	user.Email,
			"role":     user.Role,
			"exp":      expiration,
		})
		tokenString, error := token.SignedString(jwtToken)
		if error != nil {
			utils.ErrorLogger.Println(error.Error())

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		userProfile := models.UserProfile{
			Email: userService.User.Email,
			FirstName: userService.User.FirstName,
			LastName: userService.User.LastName,
			Major: userService.User.Major,
			Role: userService.User.Role,
		}

		json.NewEncoder(w).Encode(models.JwtToken{Token: tokenString, Expiration: expiration, UserProfile: userProfile})

	}
}
