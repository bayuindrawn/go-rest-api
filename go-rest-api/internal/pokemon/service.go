package pokemon

import "context"

type Service interface {
	FetchPokemonList(ctx context.Context) (*PokemonListResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) FetchPokemonList(ctx context.Context) (*PokemonListResponse, error) {
	return s.repo.FetchPokemonList(ctx)
}
