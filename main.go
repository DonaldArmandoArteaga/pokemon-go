package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-filter-pokemon-api/controllers"
	"github.com/go-filter-pokemon-api/requests"
	"github.com/go-filter-pokemon-api/services"
)

func main() {

	os.Setenv("PokemonURL", "http://3.19.60.251:3000/api/v2/pokemon")

	r := gin.Default()

	controllers.InitFilterController(
		services.Filters{
			ApiRequest: &requests.PokeApiRequest{},
		},
		r,
	)

	r.Run()

}
