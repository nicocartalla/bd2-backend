package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func getPositionTableByChampionship(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{championship_id}", controllers.GetPositionTableByChampionship).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

/* func getUserScoresByGroup(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("/{group_id}", controllers.GetUserScoresByGroup).Methods("GET")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
} */


func PositionTableRouter(r *mux.Router) {
	getPositionTableByChampionship(r)
	// getUserScoresByGroup(r)
}