package controllers

import (
	"bd2-backend/src/models"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strings"
)


func validateToken(r *http.Request) (models.User, error) {
	//obtener el token desde el header Authorization
	auth := r.Header.Get("Authorization")
	if auth != "" {
		//separar el token del string "Bearer "
		bearerToken := strings.Split(auth, " ")[1]

		// validar el token
		token, _ := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return jwtToken, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var user models.User
			mapstructure.Decode(claims, &user)
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				var user models.User
				mapstructure.Decode(claims, &user)
			}
			return user, nil
		} else {
			return models.User{}, fmt.Errorf("invalid authorization token")
		}
	} else {
		return models.User{}, fmt.Errorf("an authorization header is required")
	}
}
