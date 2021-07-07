package services

import (
	"sync"

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

	c := make(chan *models.Pokemon, 2000)
	ce := make(chan error, 2000)
	var wg sync.WaitGroup

	arr := []*models.Pokemon{}
	errs := []error{}

	for _, result := range pokemons.Results {

		wg.Add(1)

		go func(result models.Result, wg *sync.WaitGroup) {

			defer wg.Done()

			p, err := filters.ApiRequest.GetPokemonByUrlId(result.Url)

			if err != nil {
				ce <- err
			} else {
				c <- p
			}

		}(result, &wg)

	}

	wg.Wait()

	close(c)
	close(ce)

	for i := range c {
		if i.Height >= height && i.Weight >= weight {
			arr = append(arr, i)
		}
	}

	for i := range ce {
		errs = append(errs, i)
	}

	return arr, len(arr), errs, nil
}
