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

var pokeApiUrl string = "https://pokeapi.co/api/v2/pokemon"

func InitNext() {
	next = structs.Next{Offset: 0, Limit: 30}
	previous = structs.Next{Offset: 0, Limit: 30}
}

func getPokemonsExternal() {
	url := next.GetUrl(pokeApiUrl)
	response, err := http.Get(url)
	errorsHandlers.CheckNilErr(err)

	externalPokemons.SetPokemons(*response)
	next.SetInfo(externalPokemons.Next)
	files.SavePokemonsInCSV(externalPokemons)
}

func getPokemonExternal(name string, pokemon *structs.Pokemon) {
	url := pokeApiUrl + "/" + name
	response, err := http.Get(url)
	errorsHandlers.CheckNilErr(err)

	pokemon.SetPokemon(*response)
}

func getNextPokemonsExternal() {
	url := next.GetUrl(pokeApiUrl)
	updatePokemonsExternal(url)
}

func getPreviousPokemonsExternal() {
	url := previous.GetUrl(pokeApiUrl)
	updatePokemonsExternal(url)
}

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
