package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func getAllMatchesByChampionshipID(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/all").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{championship_id}", controllers.GetAllMatchesByChampionshipID).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func getAllPlayedMatchesByChampionshipID(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/played").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{championship_id}", controllers.GetAllPlayedMatchesByChampionshipID).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func getNotPlayedMatchesByChampionshipID(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/notplayed").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{championship_id}", controllers.GetNotPlayedMatchesByChampionshipID).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func getMatchResultByMatchId(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/played/id").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{match_id}", controllers.GetMatchResult).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func insertMatch(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/insert").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/", controllers.InsertMatch).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func updateMatch(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/update").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.UpdateMatch).Methods("PUT")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func deleteMatch(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/delete").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.DeleteMatch).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func insertResult(r *mux.Router) *mux.Router {
	a := r.PathPrefix("/result/insert").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.InsertResult).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func MatchRouter(r *mux.Router) {
	getAllMatchesByChampionshipID(r)
	getAllPlayedMatchesByChampionshipID(r)
	getNotPlayedMatchesByChampionshipID(r)
	getMatchResultByMatchId(r)
	insertMatch(r)
	updateMatch(r)
	deleteMatch(r)
	insertResult(r)
}
