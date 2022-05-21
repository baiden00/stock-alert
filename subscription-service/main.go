package main

import (
	"fmt"
	"net/http"

	"github.com/baiden00/subscription-service/handlers"
	"github.com/baiden00/subscription-service/structs"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func initializeRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func initializeInfra() (*structs.Cache, *structs.Datastore) {
	datastore := structs.NewDatastore("database", "addy")
	cache := structs.NewCache("cache", redis.Client{})

	return cache, datastore
}

func main() {
	fmt.Println("Hello World")
	router := initializeRouter()

	router.HandleFunc("/healthcheck", handlers.HealthCheck)
	http.ListenAndServe(":8080", router)

}
