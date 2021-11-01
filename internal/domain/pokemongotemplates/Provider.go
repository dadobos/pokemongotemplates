package pokemongotemplates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPokemons() (*[]Pokemon, error) {
	//get pokemon data from api
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/?offset=0&limit=500")

	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	// read api response
	rowData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var pokemonsData PokemonsAPIResponse
	if err := json.Unmarshal(rowData, &pokemonsData); err != nil {
		fmt.Println(err.Error())
	}

	PokemonSlice := make([]Pokemon, 0)

	for i := range pokemonsData.Results {
		resp, err := http.Get(pokemonsData.Results[i].URL)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()
		rowData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		var pokemonDetails PokemonDetails
		if err := json.Unmarshal(rowData, &pokemonDetails); err != nil {
			fmt.Println(err.Error())
		}

		var pokeMoves string
		for x := range pokemonDetails.Moves {
			if x <= 3 {
				pokeMoves += pokemonDetails.Moves[x].Move.Name + " "
			}
		}
		var pokeTypes string
		for z := range pokemonDetails.Types {
			pokeTypes += pokemonDetails.Types[z].Type.Name + " "
		}

		PokemonSlice = append(PokemonSlice, Pokemon{
			Name:   pokemonDetails.Name,
			Weight: pokemonDetails.Weight,
			Height: pokemonDetails.Height,
			Moves:  pokeMoves,
			Types:  pokeTypes,
			Image:  pokemonDetails.Sprites.Other.OfficialArtwork.FrontDefault,
		})

	}
	return &PokemonSlice, nil
}
