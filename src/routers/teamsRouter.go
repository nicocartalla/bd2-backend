package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)



func getTeams(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	// allow CORS
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetTeams).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func checkTeamExists(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/exists").Subrouter()
	// allow CORS
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{name}", controllers.CheckTeamExists).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

// Función para configurar las rutas del equipo
func TeamRouter(r *mux.Router) {
	getTeams(r)
	checkTeamExists(r)
}