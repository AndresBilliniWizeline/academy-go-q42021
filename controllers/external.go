package controllers

import (
	"challenge/api/errorsHandlers"
	"challenge/api/files"
	"fmt"
	"strings"

	structures "challenge/api/structs"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

var next structures.Next
var previous structures.Next

func InitNext() {
	next = structures.Next{Offset: 0, Limit: 30}
	previous = structures.Next{Offset: 0, Limit: 30}
}

func getPokemonsExternal() {

	pokemonsResponse, err := pokeapi.Resource("pokemon", 0, next.Limit)
	errorsHandlers.CheckNilErr(err)

	next.SetNext(pokemonsResponse.Next)

	files.SavePokemonsInCSV(pokemonsResponse)
}

func findPokemonByName(name string) (structs.Pokemon, error) {
	pokemonsResponse, err := pokeapi.Pokemon(name)

	return pokemonsResponse, err
}

func getNextPokemonsExternal() {
	pokemonsResponse, err := pokeapi.Resource("pokemon", next.Offset, next.Limit)
	errorsHandlers.CheckNilErr(err)
	next.SetNext(pokemonsResponse.Next)
	resPrevious := fmt.Sprintf("%v", pokemonsResponse.Previous)
	if !strings.Contains(resPrevious, "nil") {
		previous.SetNext(resPrevious)
	}
	files.SavePokemonsInCSV(pokemonsResponse)
}

func getPreviousPokemonsExternal() {
	pokemonsResponse, err := pokeapi.Resource("pokemon", previous.Offset, previous.Limit)
	errorsHandlers.CheckNilErr(err)
	next.SetNext(pokemonsResponse.Next)
	resPrevious := fmt.Sprintf("%v", pokemonsResponse.Previous)
	if !strings.Contains(resPrevious, "nil") {
		previous.SetNext(resPrevious)
	}
	files.SavePokemonsInCSV(pokemonsResponse)
}
