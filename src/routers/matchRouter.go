package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)



func getAllMatchResults(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetAllMatchResults).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

 func getResultsMatch(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
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

// Función para configurar las rutas del equipo
func MatchRouter(r *mux.Router) {
	getAllMatchResults(r)
	getResultsMatch(r)
	InsertMatch(r)
}