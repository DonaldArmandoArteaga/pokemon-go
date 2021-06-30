package models

type PokemonList struct {
	Count   int      `json:"count"`
	Results []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
