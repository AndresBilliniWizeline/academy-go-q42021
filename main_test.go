package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"challenge/api/structs"
)

const apiUrl = "http://localhost:4000"

type TestType struct {
	Name           string
	Pokemon        bool
	Pokemons       bool
	Concurrency    bool
	PokemonName    string
	Type           string
	Items          string
	ItemsPerWorker string
	Status         int
	Expected       int
	Err            string
}

func TestApi(t *testing.T) {
	tt := []TestType{
		{
			Name:     "get init pokemons",
			Pokemons: true,
			Status:   http.StatusOK,
			Expected: 30,
		},
		{
			Name:        "get pokemon eevee",
			Pokemon:     true,
			PokemonName: "Eevee",
			Status:      http.StatusOK,
		},
		{
			Name:        "pokemon not found",
			Pokemon:     true,
			PokemonName: "Trick",
			Status:      http.StatusOK,
			Err:         "Pokemon not found",
		},
		{
			Name:           "even pokemons with 20 items and 10 items per worker",
			Concurrency:    true,
			Type:           "even",
			Items:          "20",
			ItemsPerWorker: "10",
			Status:         http.StatusOK,
			Expected:       10,
		},
		{
			Name:           "odd pokemons with 5 items and 10 items per worker",
			Concurrency:    true,
			Type:           "odd",
			Items:          "5",
			ItemsPerWorker: "10",
			Status:         http.StatusOK,
			Expected:       5,
		},
		{
			Name:           "missing type value",
			Concurrency:    true,
			Items:          "5",
			ItemsPerWorker: "10",
			Status:         http.StatusOK,
			Expected:       5,
			Err:            "The type param is not valid",
		},
	}

	for _, tc := range tt {
		var url string
		switch true {
		case tc.Pokemon:
			url = apiUrl + "/pokemon/" + tc.PokemonName
		case tc.Pokemons:
			url = apiUrl + "/pokemons"
		case tc.Concurrency:
			url = apiUrl + "/concurrency?type=" + tc.Type + "&items=" + tc.Items + "&items_per_worker=" + tc.ItemsPerWorker
		}
		handleTest(tc, url, t)
	}
}

func handleTest(tc TestType, url string, t *testing.T) {
	srv := httptest.NewServer(initServer())
	defer srv.Close()

	res, err := http.Get(url)

	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	defer res.Body.Close()

	if tc.Err != "" {
		if tc.Pokemon {
			if res.StatusCode != http.StatusNotFound {
				t.Errorf("expected status Not Found: got %v", res.StatusCode)
			}
		} else {
			if res.StatusCode != http.StatusBadRequest {
				t.Errorf("expected status Bad Request: got %v", res.StatusCode)
			}
		}
		if msg := string(bytes.TrimSpace(b)); msg != tc.Err {
			t.Errorf("expected message %q; got %v", tc.Err, msg)
		}
		return
	}

	if res.StatusCode != tc.Status {
		t.Errorf("expected status Ok; got %v", res.StatusCode)
	}

	if tc.Pokemon {
		var pokemon structs.Pokemon
		var responseString strings.Builder
		responseString.Write(bytes.TrimSpace(b))
		content := responseString.String()
		jsonData := []byte(content)
		json.Unmarshal(jsonData, &pokemon)

		if strings.ToLower(tc.PokemonName) != pokemon.Name {
			t.Fatalf("expected pokemon to be %v; got %v", tc.PokemonName, pokemon.Name)
		}
	} else {
		var pokemons []structs.Pokemon
		var responseString strings.Builder
		responseString.Write(bytes.TrimSpace(b))
		content := responseString.String()
		jsonData := []byte(content)
		json.Unmarshal(jsonData, &pokemons)

		if len(pokemons) != tc.Expected {
			t.Fatalf("expected pokemons to be %v; got %v", tc.Expected, len(pokemons))
		}
	}
}
