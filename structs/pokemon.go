package structs

import (
	"challenge/api/errorsHandlers"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Pokemon structure to handle information in the csv file and returns responses
type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Calculates if the pokemon has an odd Id
func (p *Pokemon) Odd() bool {
	return p.Id%2 == 0
}

// Calculates if the pokemon has an even Id
func (p *Pokemon) Even() bool {
	return p.Id%2 == 1
}

// Gets a line from the csv file and set the information into a pokemon structure
func (p *Pokemon) SetInfoCSV(line []string) {
	id, _ := strconv.Atoi(line[0])
	p.Id = id
	p.Name = line[1]
	p.Url = line[2]
}

// Transform the response of a external pokemon into a pokemon structure
func (p *Pokemon) SetPokemon(response http.Response) {
	readBody, bodyErr := ioutil.ReadAll(response.Body)
	errorsHandlers.CheckNilErr(bodyErr)

	var responseString strings.Builder
	responseString.Write(readBody)
	content := responseString.String()
	jsonData := []byte(content)
	json.Unmarshal(jsonData, &p)
	p.Url = "https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(p.Id)
}
