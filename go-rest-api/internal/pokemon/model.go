package pokemon

type PokemonItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonListResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []PokemonItem `json:"results"`
}
