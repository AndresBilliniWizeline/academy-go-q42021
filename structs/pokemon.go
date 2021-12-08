package structs

import (
	"challenge/api/errorsHandlers"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Pokemon struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (p *Pokemon) Odd() bool {
	return p.Id%2 == 0
}

func (p *Pokemon) Even() bool {
	return p.Id%2 == 1
}

func (p *Pokemon) SetInfoCSV(line []string) {
	id, _ := strconv.Atoi(line[0])
	p.Id = id
	p.Name = line[1]
	p.Url = line[2]
}

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
