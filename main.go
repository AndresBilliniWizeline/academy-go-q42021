package main

import (
	"log"
	"net/http"

	"challenge/api/routers"

	"github.com/gorilla/mux"
)

func initServer() {
	router := mux.NewRouter()
	routers.Routes(router)
	log.Fatal(http.ListenAndServe(":4000", router))
}

func main() {
	initServer()
}
