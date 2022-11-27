package pokemongotemplates

type SiteLink struct {
	Title string
	Link  string
}

var NavigationLinks = []SiteLink{
	{
		Title: "All Pokemons",
		Link:  "/",
	},
	{
		Title: "With Pagination",
		Link:  "/pagination",
	},
}

type Pokemon struct {
	Name   string
	Weight int
	Height int
	Moves  string
	Types  string
	Image  string
}

type Pokemons struct {
	Count          int
	Next           string
	Previous       string
	SortedPokemons []string
	PokemonMap     map[string]*Pokemon
}

type PageData struct {
	NavigationLinks []SiteLink
	Data            Pokemons
}

type NextURL struct {
	URL string `form:"url"`
}

type PokemonMove struct {
	Name string `json:"name"`
}
type PokemonMoves struct {
	Move PokemonMove `json:"move"`
}

type PokemonType struct {
	Name string `json:"name"`
}
type PokemonTypes struct {
	Type PokemonType `json:"type"`
}

type PokemonImage struct {
	FrontDefault string `json:"front_default"`
}
type PokemonArtwork struct {
	OfficialArtwork PokemonImage `json:"official-artwork"`
}
type PokemonSprites struct {
	Other PokemonArtwork `json:"other"`
}

type PokemonDetails struct {
	Name    string         `json:"name"`
	Weight  int            `json:"weight"`
	Height  int            `json:"height"`
	Moves   []PokemonMoves `json:"moves"`
	Types   []PokemonTypes `json:"types"`
	Sprites PokemonSprites `json:"sprites"`
}

type PokemonAPIResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonsAPIResponse struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []PokemonAPIResult `json:"results"`
}
