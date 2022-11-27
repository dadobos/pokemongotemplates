package pokemongotemplates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"sync"
)

const (
	AllPokemonsURL = "https://pokeapi.co/api/v2/pokemon/?offset=0&limit=1154"
	PaginationURL  = "https://pokeapi.co/api/v2/pokemon?offset=0&limit=20"
)

func GetPokemons(url string) (Pokemons, error) {

	//get pokemon data from api
	URL := url
	if url == "" {
		URL = PaginationURL
	}

	resp, err := http.Get(URL)
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

	pokemonMap := make(map[string]*Pokemon)

	pokemonStream := make(chan Pokemon)

	var wg sync.WaitGroup
	wg.Add(len(pokemonsData.Results))

	for i := range pokemonsData.Results {
		go func(iterator int) {
			defer wg.Done()
			indPokemon := fetchIndividualPokemon(pokemonsData.Results[iterator].URL)
			pokemonStream <- indPokemon
		}(i)
	}

	for range pokemonsData.Results {
		p := <-pokemonStream
		pokemonMap[p.Name] = &p
	}

	go func() {
		wg.Wait()
		close(pokemonStream)
	}()

	sortedPokemonSlice := make([]string, len(pokemonsData.Results))
	for pokemon := range pokemonMap {
		name := pokemonMap[pokemon].Name
		sortedPokemonSlice = append(sortedPokemonSlice, name)
	}
	sort.Strings(sortedPokemonSlice)

	pokemons := Pokemons{
		Count:          pokemonsData.Count,
		Next:           pokemonsData.Next,
		Previous:       pokemonsData.Previous,
		SortedPokemons: sortedPokemonSlice,
		PokemonMap:     pokemonMap,
	}

	return pokemons, nil
}

func fetchIndividualPokemon(pokemonURL string) Pokemon {
	resp, err := http.Get(pokemonURL)
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

	return Pokemon{
		Name:   pokemonDetails.Name,
		Weight: pokemonDetails.Weight,
		Height: pokemonDetails.Height,
		Moves:  pokeMoves,
		Types:  pokeTypes,
		Image:  pokemonDetails.Sprites.Other.OfficialArtwork.FrontDefault,
	}

}
