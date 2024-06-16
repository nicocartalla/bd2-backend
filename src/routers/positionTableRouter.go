package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func getPositionTableByChampionship(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.GetPositionTableByChampionship).Methods("POST")
	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}

func PositionTableRouter(r *mux.Router) {
	getPositionTableByChampionship(r)
}
