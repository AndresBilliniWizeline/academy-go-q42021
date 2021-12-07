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
	url := next.GetNextUrl(pokeApiUrl)
	response, err := http.Get(url)
	errorsHandlers.CheckNilErr(err)

	externalPokemons.SetPokemons(*response)
	next.SetNext(externalPokemons.Next)
	files.SavePokemonsInCSV(externalPokemons)
}

func getNextPokemonsExternal() {
	url := next.GetNextUrl(pokeApiUrl)
	updatePokemonsExternal(url)
}

func getPreviousPokemonsExternal() {
	url := previous.GetNextUrl(pokeApiUrl)
	updatePokemonsExternal(url)
}

func updatePokemonsExternal(url string) {
	response, err := http.Get(url)
	errorsHandlers.CheckNilErr(err)
	externalPokemons.SetPokemons(*response)
	next.SetNext(externalPokemons.Next)
	resPrevious := fmt.Sprintf("%v", externalPokemons.Previous)
	if !strings.Contains(resPrevious, "nil") {
		previous.SetNext(resPrevious)
	}
	files.SavePokemonsInCSV(externalPokemons)
}
