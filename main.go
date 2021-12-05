package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"challenge/api/routers"
)

func initServer() {
	router := mux.NewRouter()
	routers.Routes(router)
	log.Fatal(http.ListenAndServe(":4000", router))
}

func main() {
	initServer()
}
