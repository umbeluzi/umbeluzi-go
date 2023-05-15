package umbeluzi

import (
	"context"
	"errors"
)

type ProviderListOptions struct {
	Currency *string `json:"currency"`
	Country  *string `json:"country"`
}

type ProviderListResult struct{}

type ProvidersService struct{}

func (p ProvidersService) ListProviders(context.Context, ProviderListOptions) (*ProviderListResult, error) {
	return nil, errors.New("not implemented")
}
