package infrastructure

import (
	"fmt"
	"time"
)

func GetCurrentYear() string {
	return fmt.Sprintf("%v", time.Now().Year())
}

func GetPrevPagePokemons(current int) int {
	if current < 10 {
		return current
	}

	return current - 10
}

func GetNextPagePokemons(current int) int {
	return current + 10
}
