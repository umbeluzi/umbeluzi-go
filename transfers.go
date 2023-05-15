package umbeluzi

import (
	"context"
	"errors"
)

// Transfer ...
type Transfer struct{}

// TransferListOptions ...
type TransferListOptions struct{}

// TransferCreate ...
type TransferCreate struct{}

// TransferResult ...
type TransferResult struct{}

// TransfersService ...
type TransfersService struct{}

// CreateTransfer ...
func (t TransfersService) CreateTransfer(ctx context.Context, transfer TransferCreate) (*TransferResult, error) {
	return nil, errors.New("not implemented")
}

// ConfirmTransfer ...
func (t TransfersService) ConfirmTransfer(ctx context.Context, id int64) (*TransferResult, error) {
	return nil, errors.New("not implemented")
}

// CancelTransfer ...
func (t TransfersService) CancelTransfer(ctx context.Context, id int64) (*TransferResult, error) {
	return nil, errors.New("not implemented")
}

// FetchTransfer ...
func (t TransfersService) FetchTransfer(ctx context.Context, id int64) (*TransferResult, error) {
	return nil, errors.New("not implemented")
}

// ListTransfers ...
func (t TransfersService) ListTransfers(ctx context.Context, opts *TransferListOptions) ([]TransferResult, error) {
	return nil, errors.New("not implemented")
}
