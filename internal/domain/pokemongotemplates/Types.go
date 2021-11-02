package pokemongotemplates

type SiteLink struct {
	Title string
	Link  string
}

var NavigationLinks = []SiteLink{
	{
		Title: "500 Pokemons",
		Link:  "/",
	},
	{
		Title: "Choose how many",
		Link:  "/custom",
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

type PageData struct {
	NavigationLinks []SiteLink
	Data     map[string]*Pokemon
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
	Results []PokemonAPIResult `json:"results"`
}
