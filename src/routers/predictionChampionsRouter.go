package routers

import (
	"bd2-backend/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)


func getPredictionChampionsByUserAndChampionshipID(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetPredictionChampionshipByUserAndChampionshipID).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func insertOrUpdatePredictionChampions(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/insert").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.InsertPredictionChampionship).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func PredictionChampionshipRouter(r *mux.Router) {
	getPredictionChampionsByUserAndChampionshipID(r)
	insertOrUpdatePredictionChampions(r)
}
