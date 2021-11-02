package app

import (
	"fmt"
	"net/http"

	"github.com/diaid83/pokemongotemplates/internal/domain/pokemongotemplates"
	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) pokemonPageHandler() func(*gin.Context) {
	data, err := pokemongotemplates.GetPokemons()
	if err != nil {
		fmt.Println(err.Error())
	}
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/Pokemons.gohtml", pokemongotemplates.PageData{
			NavigationLinks: pokemongotemplates.NavigationLinks,
			Data:            data,
		})
	}
}
func (s *ApplicationServer) customPokemonPageHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/CustomPokemons.gohtml", pokemongotemplates.PageData{
			NavigationLinks: pokemongotemplates.NavigationLinks,
		})
	}
}
