package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func getPredictionsByUser(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/user").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{user_id}", controllers.GetPredictionsByUser).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func createPrediction(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/create").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.InsertPrediction).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}


func PredictionRouter(r *mux.Router) {
	getPredictionsByUser(r)
	createPrediction(r)
}