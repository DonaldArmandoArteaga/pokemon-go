package controllers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-filter-pokemon-api/services"
)

type FilterController struct {
	filter services.Filters
	gin    *gin.Engine
}

func InitFilterController(filter services.Filters, gin *gin.Engine) {
	fc := FilterController{
		filter: filter,
		gin:    gin,
	}

	fc.gin.GET("/pokemons", fc.getByWeightAndHeight)

}

func (fc *FilterController) getByWeightAndHeight(c *gin.Context) {

	w, errw := strconv.Atoi(c.Query("weight"))
	h, errh := strconv.Atoi(c.Query("height"))

	errarr := []string{}

	if errw != nil {
		errarr = append(errarr, "weight must be a integer")
	}

	if errh != nil {
		errarr = append(errarr, "height must be a integer")
	}

	if len(errarr) > 0 {
		c.JSON(500, gin.H{
			"error": strings.Join(errarr, ", "),
		})
		return
	}

	values, count, errors, err := fc.filter.WeightAndHeight(h, w)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"count":  count,
		"values": values,
		"errors": errors,
	})

}
