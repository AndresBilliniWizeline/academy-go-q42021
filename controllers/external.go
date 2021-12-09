package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"challenge/api/errorsHandlers"
	"challenge/api/files"
	"challenge/api/structs"
)

var next structs.Next
var previous structs.Next
var externalPokemons structs.ExternalPokemon

// External API
const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon"

// Set values of next
func InitNext() {
	next = structs.Next{Offset: 0, Limit: 30}
	previous = structs.Next{Offset: 0, Limit: 30}
}

// Gets an array of pokemons from an external api
func getPokemonsExternal() {
	url := next.GetUrl(pokeApiUrl)
	response, err := http.Get(url)
	errorsHandlers.CheckNilErr(err)

	externalPokemons.SetPokemons(*response)
	next.SetInfo(externalPokemons.Next)
	files.SavePokemonsInCSV(externalPokemons)
}

// Gets the pokemon that you were searching for from an external api
func getPokemonExternal(name string, pokemon *structs.Pokemon) {
	url := pokeApiUrl + "/" + name
	response, err := http.Get(url)
	errorsHandlers.CheckNilErr(err)

	pokemon.SetPokemon(*response)
}

// Gets the next array of pokemons from an external api
func getNextPokemonsExternal() {
	url := next.GetUrl(pokeApiUrl)
	updatePokemonsExternal(url)
}

// Gets the previous array of pokemons from an external api
func getPreviousPokemonsExternal() {
	url := previous.GetUrl(pokeApiUrl)
	updatePokemonsExternal(url)
}

// updates the csv file
func updatePokemonsExternal(url string) {
	response, err := http.Get(url)
	errorsHandlers.CheckNilErr(err)
	externalPokemons.SetPokemons(*response)
	next.SetInfo(externalPokemons.Next)
	resPrevious := fmt.Sprintf("%v", externalPokemons.Previous)
	if !strings.Contains(resPrevious, "nil") {
		previous.SetInfo(resPrevious)
	}
	files.SavePokemonsInCSV(externalPokemons)
}
