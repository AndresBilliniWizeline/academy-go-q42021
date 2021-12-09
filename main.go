package main

import (
	"log"
	"net/http"

	"challenge/api/routers"

	"github.com/gorilla/mux"
)

func initServer() *mux.Router {
	router := mux.NewRouter()
	routers.Routes(router)
	return router
}

func main() {
	router := initServer()
	log.Fatal(http.ListenAndServe(":4000", router))
}
