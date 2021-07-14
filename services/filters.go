package services

import (
	"github.com/go-filter-pokemon-api/models"
	"github.com/go-filter-pokemon-api/requests"
)

type PokemonFilters interface {
	WeightAndHeight(height int, weight int) ([]*models.Pokemon, int, []error, error)
}

type Filters struct {
	ApiRequest requests.PokemonRequest
}

func (filters *Filters) WeightAndHeight(height int, weight int) ([]*models.Pokemon, int, []error, error) {

	pokemons, err := filters.ApiRequest.GetAllPokemon()

	if err != nil {
		return nil, 0, nil, err
	}

	arr := []*models.Pokemon{}
	errs := []error{}

	for _, result := range pokemons.Results {

		p, err := filters.ApiRequest.GetPokemonByUrlId(result.Url)

		if err != nil {
			errs = append(errs, err)
		} else {
			arr = append(arr, p)
		}

	}

	return arr, len(arr), errs, nil
}
