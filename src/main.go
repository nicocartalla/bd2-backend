package main

import (
	"bd2-backend/src/config"
	"bd2-backend/src/routers"
	"bd2-backend/src/services"
	"log"
	"net/http"
	"time"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	router := routers.Routers()
	log.Println("Server started at address", config.ServerAddress+":8080")

	go func() {
		log.Fatal(http.ListenAndServe(config.ServerAddress+":8080", router))
	}()


	go func() {
		services.SendEmailToUsersWhoHaveNotMadePredictions()
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			services.SendEmailToUsersWhoHaveNotMadePredictions()
		}
	}()
	
	select {}
}
