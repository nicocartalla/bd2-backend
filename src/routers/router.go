package routers

import (
	"bd2-backend/src/utils"
	"bd2-backend/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)


func Routers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	enableCORS(r)

	//api version 1
	v1 := r.PathPrefix("/api/v1").Subrouter()
	ping := v1.PathPrefix("/ping").Subrouter()
	auth := v1.PathPrefix("/authenticate").Subrouter()
	signup := v1.PathPrefix("/signup").Subrouter()
	team := v1.PathPrefix("/team").Subrouter()
	match := v1.PathPrefix("/match").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	utils.InfoLogger.Println("CORS enabled")

	PingRouter(ping)
	utils.InfoLogger.Println("Ping router enabled at /api/v1/ping")
	AuthRouter(auth)
	utils.InfoLogger.Println("Auth router enabled at /api/v1/authenticate")
	SignUpRouter(signup)
	utils.InfoLogger.Println("User router enabled at /api/v1/signup")
	TeamRouter(team)
	utils.InfoLogger.Println("User router enabled at /api/v1/team")
	MatchRouter(match)
	utils.InfoLogger.Println("User router enabled at /api/v1/match")
	
	return r
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses.Exception{Message: "path not found"})
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses.Exception{Message: "method not allowed"})
	}
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
