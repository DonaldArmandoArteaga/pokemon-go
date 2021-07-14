package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-filter-pokemon-api/models"
)

type PokemonRequest interface {
	GetAllPokemon() (*models.PokemonList, error)
	GetPokemonByUrlId(url string) (*models.Pokemon, error)
}

type PokeApiRequest struct {
}

func (*PokeApiRequest) GetPokemonByUrlId(url string) (*models.Pokemon, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var pokemon *models.Pokemon
	json.Unmarshal(body, &pokemon)
	fmt.Println(pokemon)
	return pokemon, nil
}

func (*PokeApiRequest) GetAllPokemon() (*models.PokemonList, error) {

	resp, err := http.Get(os.Getenv("PokemonURL"))

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var pokemons *models.PokemonList
	json.Unmarshal(body, &pokemons)

	return pokemons, nil

}
