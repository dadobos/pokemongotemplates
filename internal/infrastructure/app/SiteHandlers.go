package app

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dadobos/pokemongotemplates/internal/domain/pokemongotemplates"
	"github.com/gin-gonic/gin"
)

func (s *ApplicationServer) pageHandler() func(*gin.Context) {
	data, err := pokemongotemplates.GetPokemons(pokemongotemplates.AllPokemonsURL)
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

func (s *ApplicationServer) paginationPageHandler() func(*gin.Context) {
	return func(c *gin.Context) {

		requestURL := pokemongotemplates.PaginationURL
		if newURL := c.Request.URL.Query(); newURL.Has("slug") {
			requestURL = newURL["slug"][0]
		}

		data, err := pokemongotemplates.GetPokemons(requestURL)
		if err != nil {
			fmt.Println(err.Error())
		}

		c.HTML(http.StatusOK, "views/PaginationPokemons.gohtml", pokemongotemplates.PageData{
			NavigationLinks: pokemongotemplates.NavigationLinks,
			Data:            data,
		})
	}
}


func (s *ApplicationServer) slugSendPaginationRequestHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestData pokemongotemplates.NextURL
		if err := c.Bind(&requestData); err != nil {
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.Redirect(http.StatusFound, fmt.Sprintf("/pagination/?slug=%s", url.QueryEscape(requestData.URL)))
	}
}
