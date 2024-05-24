package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)



func getAllMatchResults(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/result").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetAllMatchResults).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func getResultsMatch(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/result/id").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{match_id}", controllers.GetMatchResult).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func InsertMatch(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/create").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.InsertMatch).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func UpdateMatch(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/update").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.UpdateMatch).Methods("PUT")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func DeleteMatch(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/delete").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.DeleteMatch).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func InsertResult(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/result/insert").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.InsertResult).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func GetMatchesNotPlayedYet(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/notplayed").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetMatchesNotPlayedYet).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

// Función para configurar las rutas del equipo
func MatchRouter(r *mux.Router) {
	getAllMatchResults(r)
	getResultsMatch(r)
	InsertMatch(r)
	UpdateMatch(r)
	DeleteMatch(r)
	InsertResult(r)
	GetMatchesNotPlayedYet(r)
}