package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func getTeams(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetTeams).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func getTeamsByChampionshipID(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/championship").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{id}", controllers.GetTeamsByChampionshipID).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func getTeamsByID(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/{id}").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetTeamByID).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func checkTeamExists(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/exists").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{name}", controllers.CheckTeamExists).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func TeamRouter(r *mux.Router) {
	getTeams(r)
	getTeamsByID(r)
	checkTeamExists(r)
	getTeamsByChampionshipID(r)
}
