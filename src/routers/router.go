package routers

import (
	"bd2-backend/src/utils"
	"bd2-backend/src/responses"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)


func Routers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	enableCORS(r)

	//auth
	auth := r.PathPrefix("/api/auth").Subrouter()
	signIn := auth.PathPrefix("/signin").Subrouter()
	signUp := auth.PathPrefix("/signup").Subrouter()

	//api version 1
	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.Use(middlewareJwt)
	ping := v1.PathPrefix("/ping").Subrouter()
	team := v1.PathPrefix("/team").Subrouter()
	match := v1.PathPrefix("/match").Subrouter()
	prediction := v1.PathPrefix("/prediction").Subrouter()
	positiontable := v1.PathPrefix("/positiontable").Subrouter()
	championship := v1.PathPrefix("/championship").Subrouter()
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	utils.InfoLogger.Println("CORS enabled")

	
	SignInRouter(signIn)
	utils.InfoLogger.Println("Auth router enabled at /api/auth/signin")
	SignUpRouter(signUp)
	utils.InfoLogger.Println("User router enabled at /api/auth/signup")

	PingRouter(ping)
	utils.InfoLogger.Println("Ping router enabled at /api/v1/ping")
	TeamRouter(team)
	utils.InfoLogger.Println("User router enabled at /api/v1/team")
	MatchRouter(match)
	utils.InfoLogger.Println("User router enabled at /api/v1/match")
	PredictionRouter(prediction)
	utils.InfoLogger.Println("User router enabled at /api/v1/prediction")
	PositionTableRouter(positiontable)
	utils.InfoLogger.Println("User router enabled at /api/v1/positiontable")
	ChampionshipRouter(championship)
	utils.InfoLogger.Println("User router enabled at /api/v1/championship")
	
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


func middlewareJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  auth_header := r.Header.Get("Authorization")
	  if !strings.HasPrefix(auth_header, "Bearer") {
		  http.Error(w, "Not Authorized", http.StatusUnauthorized)
		  return
	  }
	  
	  tokenString := strings.TrimPrefix(auth_header, "Bearer ")
	  
	  claims, err := utils.GetClaimsFromToken(tokenString)
	  if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
  }
	  r = r.WithContext(utils.SetJWTClaimsContext(r.Context(), claims))
	  next.ServeHTTP(w, r)
  })
  }