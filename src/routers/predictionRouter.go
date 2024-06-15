package routers

import (
	"bd2-backend/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func getPredictionsByUser(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/user").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{document_id}", controllers.GetPredictionsByUser).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func getPredictionsByUserAndChampionshipID(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetPredictionsByUserAndChampionshipID).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func insertOrUpdatePrediction(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/insert").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.InsertPrediction).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func PredictionRouter(r *mux.Router) {
	getPredictionsByUser(r)
	getPredictionsByUserAndChampionshipID(r)
	insertOrUpdatePrediction(r)
}
