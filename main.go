package main

import (
	"fmt"
	"os"

	"github.com/go-filter-pokemon-api/requests"
	"github.com/go-filter-pokemon-api/services"
)

func main() {

	os.Setenv("PokemonURL", "http://18.119.14.15:3000/api/v2/pokemon")

	/*
		r := gin.Default()

		controllers.InitFilterController(
			services.Filters{
				ApiRequest: &requests.PokeApiRequest{},
			},
			r,
		)

		r.Run()

	*/

	a := services.Filters{
		ApiRequest: &requests.PokeApiRequest{},
	}

	v, _, _, _ := a.WeightAndHeight(100, 100)

	fmt.Println(v)

}
