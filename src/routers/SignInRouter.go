package routers

import (
	"bd2-backend/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SignInRouter(r *mux.Router) *mux.Router {
	a := r.PathPrefix("").Subrouter()
	// allow CORS
	a.Use(mux.CORSMethodMiddleware(a))
	a.HandleFunc("", controllers.SignIn).Methods("POST")

	a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
	return a
}
