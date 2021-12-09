package structs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"challenge/api/errorsHandlers"
)

// Structure for external API calls
type ExternalPokemon struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Pokemon   `json:"results"`
}

// Set the external response of pokemons in a json format
func (p *ExternalPokemon) SetPokemons(response http.Response) {
	readBody, bodyErr := ioutil.ReadAll(response.Body)
	errorsHandlers.CheckNilErr(bodyErr)
	defer response.Body.Close()

	var responseString strings.Builder
	responseString.Write(readBody)
	content := responseString.String()
	jsonData := []byte(content)
	json.Unmarshal(jsonData, &p)
}
