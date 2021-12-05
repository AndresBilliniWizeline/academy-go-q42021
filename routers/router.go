package routers

import (
	"challenge/api/controllers"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	controllers.InitNext()
	router.HandleFunc("/", controllers.HomePage).Methods("GET")
	router.HandleFunc("/pokemons", controllers.GetPokemons).Methods("GET")
	router.HandleFunc("/pokemon/{name}", controllers.GetPokemon).Methods("GET")
	router.HandleFunc("/pokemons/next", controllers.GetNextPokemons).Methods("GET")
	router.HandleFunc("/pokemons/previous", controllers.GetPreviousPokemons).Methods("GET")
	router.HandleFunc("/concurrency", controllers.ConcurrencyGetPokemons).Methods("GET")
}
