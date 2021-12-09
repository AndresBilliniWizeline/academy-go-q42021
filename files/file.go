package files

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"challenge/api/errorsHandlers"
	"challenge/api/structs"
)

const pokedex = "./db/api.csv"

func GetPokemonInfoCSV() []structs.Pokemon {
	csvFile, err := os.Open(pokedex)
	errorsHandlers.CheckFileErr(err)
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	errorsHandlers.CheckNilErr(err)

	pokemons := []structs.Pokemon{}

	for _, line := range csvLines {
		pokemon := structs.Pokemon{}
		pokemon.SetInfoCSV(line)
		pokemons = append(pokemons, pokemon)
	}
	return pokemons
}

func SavePokemonsInCSV(pokemons structs.ExternalPokemon) {
	csvFile, err := os.Create(pokedex)
	errorsHandlers.CheckFileErr(err)

	csvwriter := csv.NewWriter(csvFile)

	for index, pokemon := range pokemons.Results {
		var row []string
		row = append(row, strconv.Itoa(index+1))
		row = append(row, pokemon.Name)
		row = append(row, pokemon.Url)
		csvwriter.Write(row)
	}
	fmt.Println("Save in csvFile")
	csvwriter.Flush()

	csvFile.Close()
}
