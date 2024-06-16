package controllers

// import (
// 	"bd2-backend/src/models"
// 	"fmt"
// 	"github.com/golang-jwt/jwt/v4"
// 	"github.com/mitchellh/mapstructure"
// 	"net/http"
// 	"strings"
// )

// func validateToken(r *http.Request) (jwt.Claims, error) {
// 	auth := r.Header.Get("Authorization")
// 	if auth != "" {
// 		bearerToken := strings.Split(auth, " ")[1]

// 		token, _ := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("there was an error")
// 			}
// 			return jwtToken, nil
// 		})
// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

// 			}
// 			return user, nil
// 		} else {
// 			return models.User{}, fmt.Errorf("invalid authorization token")
// 		}
// 	} else {
// 		return models.User{}, fmt.Errorf("an authorization header is required")
// 	}
// }
