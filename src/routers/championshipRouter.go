
package routers

import (
    "bd2-backend/src/controllers"
    "github.com/gorilla/mux"
    "net/http"
)

func getAllChampionships(r *mux.Router) *mux.Router {
    a := r.PathPrefix("").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("", controllers.GetAllChampionships).Methods("GET")
    a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
    return a
}

func getChampionshipByID(r *mux.Router) *mux.Router {
    a := r.PathPrefix("/{id}").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("", controllers.GetChampionshipByID).Methods("GET")
    a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
    return a
}

func createChampionship(r *mux.Router) *mux.Router {
    a := r.PathPrefix("").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("", controllers.CreateChampionship).Methods("POST")
    a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
    return a
}

func updateChampionship(r *mux.Router) *mux.Router {
    a := r.PathPrefix("/{id}").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("", controllers.UpdateChampionship).Methods("PUT")
    a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
    return a
}

func deleteChampionship(r *mux.Router) *mux.Router {
    a := r.PathPrefix("/{id}").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("", controllers.DeleteChampionship).Methods("DELETE")
    a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
    return a
}

func setChampionshipChampions(r *mux.Router) *mux.Router {
    a := r.PathPrefix("/setchampions").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("", controllers.SetChampionshipChampions).Methods("POST")
    a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
    return a
}

func ChampionshipRouter(r *mux.Router) {
    getAllChampionships(r)
    getChampionshipByID(r)
    createChampionship(r)
    updateChampionship(r)
    deleteChampionship(r)
    setChampionshipChampions(r)
}
