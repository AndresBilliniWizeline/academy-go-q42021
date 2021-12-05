package files

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"challenge/api/errorsHandlers"
	structures "challenge/api/structs"

	"github.com/mtslzr/pokeapi-go/structs"
)

const pokedex = "./files/api.csv"

func GetPokemonInfoCSV() []structures.Pokemon {
	csvFile, err := os.Open(pokedex)
	errorsHandlers.CheckFileErr(err)
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	errorsHandlers.CheckNilErr(err)

	pokemons := []structures.Pokemon{}

	for _, line := range csvLines {
		pokemon := structures.Pokemon{}
		pokemon.SetInfoCSV(line)
		pokemons = append(pokemons, pokemon)
	}
	return pokemons
}

func SavePokemonsInCSV(pokemons structs.Resource) {
	csvFile, err := os.Create(pokedex)
	errorsHandlers.CheckFileErr(err)

	csvwriter := csv.NewWriter(csvFile)

	for index, pokemon := range pokemons.Results {
		var row []string
		row = append(row, strconv.Itoa(index+1))
		row = append(row, pokemon.Name)
		row = append(row, pokemon.URL)
		csvwriter.Write(row)
	}
	fmt.Println("Save in csvFile")
	csvwriter.Flush()

	csvFile.Close()
}
