package structs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"challenge/api/errorsHandlers"
)

type ExternalPokemon struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous *Previous `json:"previous"`
	Results  []Pokemon `json:"results"`
}

type Previous struct {
	Previous string `json:"previous"`
}

func (p *ExternalPokemon) SetPokemons(response http.Response) {
	readBody, bodyErr := ioutil.ReadAll(response.Body)
	errorsHandlers.CheckNilErr(bodyErr)
	body := string(readBody)

	fmt.Println("response", body)
	var responseString strings.Builder
	responseString.Write(readBody)
	content := responseString.String()
	jsonData := []byte(content)
	json.Unmarshal(jsonData, &p)
}
