package OracleDrex

import (
	"context"
	"log"
	pkg "service/pkg/smart-contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TransferRepository provides methods for interacting with the Transfer event
type TransferRepository struct {
	contract *pkg.OracleDrexerc3643
	client   *ethclient.Client
}

// NewTransferRepository creates a new TransferRepository
func NewTransferRepository(clientURL string, contractAddress common.Address) (*TransferRepository, error) {
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		return nil, err
	}

	contract, err := pkg.NewOracleDrexerc3643(contractAddress, client)
	if err != nil {
		return nil, err
	}

	return &TransferRepository{
		contract: contract,
		client:   client,
	}, nil
}

// ListenToTransfer listens to the Transfer event and calls the provided callback function
func (tr *TransferRepository) ListenToTransfer(ctx context.Context, callback func(event pkg.OracleDrexerc3643Transfer) error) error {
	opts := &bind.WatchOpts{Context: ctx}

	events := make(chan *pkg.OracleDrexerc3643Transfer)

	sub, err := tr.contract.WatchTransfer(opts, events)
	if err != nil {
		return err
	}

	go func() {
		defer sub.Unsubscribe()

		for {
			select {
			case event := <-events:
				// Convert event data to a more usable format
				transferEvent := pkg.OracleDrexerc3643Transfer{
					From:   event.From.Hex(),
					To:     event.To.Hex(),
					Amount: event.Amount,
				}

				// Call the callback function
				err := callback(transferEvent)
				if err != nil {
					// Handle error from the callback, e.g., log the error
					log.Printf("Error in callback: %v", err)
				}
			case err := <-sub.Err():
				// Handle subscription error
				log.Printf("Subscription error: %v", err)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
