package services_test

import (
	"errors"
	"testing"

	"github.com/go-filter-pokemon-api/models"
	"github.com/go-filter-pokemon-api/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPokemonRequest struct {
	mock.Mock
}

func (m *mockPokemonRequest) GetAllPokemon() (*models.PokemonList, error) {

	args := m.Called()
	return args.Get(0).(*models.PokemonList), args.Error(1)

}

func (m *mockPokemonRequest) GetPokemonByUrlId(url string) (*models.Pokemon, error) {

	args := m.Called(url)
	return args.Get(0).(*models.Pokemon), args.Error(1)
}

func TestWeightAndHeightFilters(t *testing.T) {

	filterMock := mockPokemonRequest{}

	filterMock.On("GetAllPokemon").Return(
		&models.PokemonList{
			Count: 1,
			Results: []models.Result{
				{
					Name: "bulbasaur",
					Url:  "https://dominio/api/v2/pokemon/1/",
				},
			},
		},
		nil,
	)

	filterMock.On("GetPokemonByUrlId", mock.AnythingOfType("string")).Return(
		&models.Pokemon{
			Height: 100,
			Weight: 100,
			Name:   "bulbasaur",
			Id:     1,
		},
		nil,
	)

	type divtest struct {
		Name          string
		A             int
		B             int
		ExpectedArray []*models.Pokemon
		ExpectedCount int
	}

	tests := []divtest{
		{
			Name: "Weigh 100 and Height 100",
			A:    100,
			B:    100,
			ExpectedArray: []*models.Pokemon{
				{
					Height: 100,
					Weight: 100,
					Name:   "bulbasaur",
					Id:     1,
				},
			},
			ExpectedCount: 1,
		},
		{
			Name: "Weigh 100 and Height 10",
			A:    100,
			B:    10,
			ExpectedArray: []*models.Pokemon{
				{
					Height: 100,
					Weight: 100,
					Name:   "bulbasaur",
					Id:     1,
				},
			},
			ExpectedCount: 1,
		},
		{
			Name:          "Weigh 1000 and Height 1000",
			A:             1000,
			B:             1000,
			ExpectedArray: []*models.Pokemon{},
			ExpectedCount: 0,
		},
		{
			Name:          "Weigh 1000 and Height 100",
			A:             1000,
			B:             100,
			ExpectedArray: []*models.Pokemon{},
			ExpectedCount: 0,
		},
	}

	filter := services.Filters{
		ApiRequest: &filterMock,
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result, count, _, _ := filter.WeightAndHeight(test.A, test.B)

			assert.ElementsMatch(t, result, test.ExpectedArray)
			assert.Equal(t, count, test.ExpectedCount)

		})
	}

}

func TestGetAllPokemonError(t *testing.T) {

	filterMock := mockPokemonRequest{}

	filterMock.On("GetAllPokemon").Return(&models.PokemonList{}, errors.New("Error test"))

	filter := services.Filters{
		ApiRequest: &filterMock,
	}

	t.Run("errorTest", func(t *testing.T) {
		_, _, _, err := filter.WeightAndHeight(100, 100)
		assert.EqualError(t, err, "Error test")
	})

}

func TestGetPokemonByUrlIdError(t *testing.T) {
	pokemonReq := mockPokemonRequest{}
	pokemonReq.On("GetAllPokemon").Return(
		&models.PokemonList{
			Count: 1,
			Results: []models.Result{
				{
					Name: "Charmander",
					Url:  "https://dominio/pokemons",
				},
			},
		},
		nil,
	)
	pokemonReq.On("GetPokemonByUrlId", mock.AnythingOfType("string")).Return(
		&models.Pokemon{},
		errors.New("Error Url test"),
	)
	filter := services.Filters{
		ApiRequest: &pokemonReq,
	}
	t.Run("errorTestURL", func(t *testing.T) {
		_, _, err, _ := filter.WeightAndHeight(100, 100)
		assert.ElementsMatch(t, err, []error{
			errors.New("Error Url test"),
		})
	})
}
