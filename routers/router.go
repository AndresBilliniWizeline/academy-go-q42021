package routers

import (
	"net/http"

	"challenge/api/controllers"

	"github.com/gorilla/mux"
)

const get = http.MethodGet

func Routes(router *mux.Router) {
	controllers.InitNext()
	router.HandleFunc("/", controllers.HomePage).Methods(get)
	router.HandleFunc("/pokemons", controllers.GetPokemons).Methods(get)
	router.HandleFunc("/pokemon/{name}", controllers.GetPokemon).Methods(get)
	router.HandleFunc("/pokemons/next", controllers.GetNextPokemons).Methods(get)
	router.HandleFunc("/pokemons/previous", controllers.GetPreviousPokemons).Methods(get)
	router.HandleFunc("/concurrency", controllers.ConcurrencyGetPokemons).Methods(get)
}
