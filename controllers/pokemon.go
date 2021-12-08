package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"challenge/api/files"
	"challenge/api/structs"

	"github.com/gorilla/mux"
)

var pokemons []structs.Pokemon
var concurrencyPokemons []structs.Pokemon
var wg sync.WaitGroup
var mut sync.Mutex

func initPokemons() {
	fmt.Println("Set all pokemons")
	pokemons = nil
	pokemons = append(pokemons, files.GetPokemonInfoCSV()...)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my Pokemon API with golang</h1>"))
}

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	getPokemonsExternal()
	initPokemons()
	fmt.Println("Get pokemons")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get pokemon")
	w.Header().Set("Content-Type", "application/json")
	// grab name from request
	params := mux.Vars(r)
	name := strings.ToLower(params["name"])

	pokemon := &structs.Pokemon{}
	getPokemonExternal(name, pokemon)

	if pokemon == nil || pokemon.Id == 0 {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(pokemon)
}

func GetNextPokemons(w http.ResponseWriter, r *http.Request) {
	getNextPokemonsExternal()
	initPokemons()
	fmt.Println("Get next pokemons")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func GetPreviousPokemons(w http.ResponseWriter, r *http.Request) {
	getPreviousPokemonsExternal()
	initPokemons()
	fmt.Println("Get previous pokemons")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func ConcurrencyGetPokemons(w http.ResponseWriter, r *http.Request) {
	initPokemons()
	fmt.Println("Get Concurrency pokemons")
	w.Header().Set("Content-Type", "application/json")

	var query structs.Query
	rawQuery := r.URL.Query()
	query.SetValues(rawQuery)
	fmt.Println(query)

	queryError, multiple := query.HandleError()

	if multiple > 0 {
		query.SendErrorMessage(w, queryError, multiple)
		return
	}

	concurrencyPokemons = nil
	for _, pokemon := range pokemons {
		go getEvenOrOdd(query, pokemon)
		wg.Add(1)
	}
	wg.Wait()
	json.NewEncoder(w).Encode(concurrencyPokemons)
}

func getEvenOrOdd(
	query structs.Query,
	pokemon structs.Pokemon,
) {
	defer wg.Done()
	if len(concurrencyPokemons) < query.Items {
		switch query.Type {
		case "odd":
			if pokemon.Odd() && len(concurrencyPokemons) < query.ItemsPerWorker {
				mut.Lock()
				concurrencyPokemons = append(concurrencyPokemons, pokemon)
				mut.Unlock()
			}
		case "even":
			if pokemon.Even() && len(concurrencyPokemons) < query.ItemsPerWorker {
				mut.Lock()
				concurrencyPokemons = append(concurrencyPokemons, pokemon)
				mut.Unlock()
			}
		}
	}

}
