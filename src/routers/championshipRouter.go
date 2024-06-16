
package routers

import (
    "bd2-backend/src/controllers"
    "github.com/gorilla/mux"
    "net/http"
)

func getAllChampionships(r *mux.Router) *mux.Router {
    a := r.PathPrefix("/all").Subrouter()
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
    a := r.PathPrefix("/create").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("", controllers.CreateChampionship).Methods("POST")
    a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
    return a
}

func updateChampionship(r *mux.Router) *mux.Router {
    a := r.PathPrefix("/update").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("/{id}", controllers.UpdateChampionship).Methods("PUT")
    a.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowed)
    return a
}

func deleteChampionship(r *mux.Router) *mux.Router {
    a := r.PathPrefix("/delete").Subrouter()
    a.Use(mux.CORSMethodMiddleware(a))
    a.HandleFunc("/{id}", controllers.DeleteChampionship).Methods("DELETE")
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
    createChampionship(r)
    setChampionshipChampions(r)
    getAllChampionships(r)
    updateChampionship(r)
    deleteChampionship(r)
    getChampionshipByID(r)
}
