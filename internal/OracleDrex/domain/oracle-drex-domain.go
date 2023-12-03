package OracleDrex

import (
	"context"
	"math/big"
)

// EventTransfer represents the Transfer event data
type EventTransfer struct {
	From   string
	To     string
	Amount *big.Int
}

// TransferListener interface defines the methods for listening to Transfer events
type TransferListener interface {
	ListenToTransfer(ctx context.Context, callback func(event EventTransfer) error) error
}
