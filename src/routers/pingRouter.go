package routers

import (
	"bd2-backend/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func PingRouter(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	// allow CORS
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.Ping).Methods("Get")

	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}
