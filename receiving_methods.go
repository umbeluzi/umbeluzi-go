package umbeluzi

import (
	"context"
	"errors"
)

// ReceivingMethodCheckResult ...
type ReceivingMethodCheckResult struct{}

// ReceivingMethodFetchResult ...
type ReceivingMethodFetchResult struct{}

// ReceivingMethodFetch ...
type ReceivingMethodFetch struct{}

// ReceivingMethodCheck ...
type ReceivingMethodCheck struct{}

type ReceivingMethodsService struct{}

// CheckReceivingMethod ...
func (a ReceivingMethodsService) CheckReceivingMethod(context.Context, ReceivingMethodCheck) (*ReceivingMethodCheckResult, error) {
	return nil, errors.New("not implemented")
}

// FetchReceivingMethod ...
func (a ReceivingMethodsService) FetchReceivingMethod(context.Context, ReceivingMethodFetch) (*ReceivingMethodFetchResult, error) {
	return nil, errors.New("not implemented")
}
