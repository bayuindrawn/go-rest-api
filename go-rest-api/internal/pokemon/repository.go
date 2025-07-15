package pokemon

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Repository interface {
	FetchPokemonList(ctx context.Context) (*PokemonListResponse, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FetchPokemonList(ctx context.Context) (*PokemonListResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://pokeapi.co/api/v2/pokemon/?limit=10", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch from PokeAPI")
	}

	var data PokemonListResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
