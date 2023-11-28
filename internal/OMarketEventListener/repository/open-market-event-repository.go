package OMarketEventListener

import (
	"context"
	"errors"
	"log"
	"math/big"
	pkg "service/pkg/smart-contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractRepository struct {
	contract *pkg.OpenMarket
	client   *ethclient.Client
}

func NewContractRepository(clientURL string, contractAddress common.Address) (*ContractRepository, error) {
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	log.Printf("Ethereum client connected successfully")

	contract, err := pkg.NewOpenMarket(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to bind to deployed instance of contract: %v", err)
		return nil, err
	}

	return &ContractRepository{contract: contract, client: client}, nil
}

func (cr *ContractRepository) GetCurrentBlockNumber() (uint64, error) {
	if cr.client == nil {
		log.Println("Ethereum client not initialized")
		return 0, errors.New("ethereum client not initialized")
	}

	header, err := cr.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Printf("Error getting current block number: %v", err)
		return 0, err
	}

	return header.Number.Uint64(), nil
}

func (cr *ContractRepository) ListenToPublicOrderCreated(startBlock uint64) {
	opts := &bind.FilterOpts{Start: startBlock, End: nil, Context: context.Background()}
	it, err := cr.contract.FilterPublicOrderCreated(opts, nil)
	if err != nil {
		log.Fatalf("Failed to filter event logs: %v", err)
	}

	for it.Next() {
		event := it.Event
		log.Printf("PublicOrderCreated event received: TokenID=%d", event.TokenId)
	}
	if it.Error() != nil {
		log.Fatalf("Iterator error: %v", it.Error())
	}
}

func (cr *ContractRepository) ListenToPrimarySale(ctx context.Context, startBlock uint64) (<-chan *pkg.OpenMarketPrimarySale, error) {
	events := make(chan *pkg.OpenMarketPrimarySale)

	go func() {
		defer close(events)

		opts := &bind.FilterOpts{
			Start:   startBlock,
			End:     nil,
			Context: ctx,
		}

		it, err := cr.contract.FilterPrimarySale(opts, nil, nil, nil)
		if err != nil {
			log.Printf("Failed to filter event logs: %v", err)
			return
		}

		for it.Next() {
			select {
			case <-ctx.Done():
				return
			case events <- it.Event:
			}
		}

		if it.Error() != nil {
			log.Printf("Iterator error: %v", it.Error())
		}
	}()

	return events, nil
}

func (cr *ContractRepository) ListenSecondaryForSale(ctx context.Context, startBlock uint64) (<-chan *pkg.OpenMarketSecondaryForSale, error) {
	events := make(chan *pkg.OpenMarketSecondaryForSale)

	go func() {
		defer close(events)

		opts := &bind.FilterOpts{
			Start:   startBlock,
			End:     nil,
			Context: ctx,
		}

		var sellers []common.Address // Filtro vazio para endereços dos vendedores
		var tokenIds []*big.Int      // Filtro vazio para IDs de token
		var units []*big.Int         // Filtro vazio para unidades

		it, err := cr.contract.FilterSecondaryForSale(opts, sellers, tokenIds, units)
		if err != nil {
			log.Printf("Failed to filter event logs: %v", err)
			return
		}

		for it.Next() {
			select {
			case <-ctx.Done():
				return
			case events <- it.Event:
			}
		}

		if it.Error() != nil {
			log.Printf("Iterator error: %v", it.Error())
		}
	}()

	return events, nil
}

func (cr *ContractRepository) ListenSecondarySold(ctx context.Context, startBlock uint64) (<-chan *pkg.OpenMarketSecondarySold, error) {
	events := make(chan *pkg.OpenMarketSecondarySold)

	go func() {
		defer close(events)

		opts := &bind.FilterOpts{
			Start:   startBlock,
			End:     nil,
			Context: ctx,
		}

		it, err := cr.contract.FilterSecondarySold(opts, nil, nil, nil)
		if err != nil {
			log.Printf("Failed to filter event logs: %v", err)
			return
		}

		for it.Next() {
			select {
			case <-ctx.Done():
				return
			case events <- it.Event:
			}
		}

		if it.Error() != nil {
			log.Printf("Iterator error: %v", it.Error())
		}
	}()

	return events, nil
}
