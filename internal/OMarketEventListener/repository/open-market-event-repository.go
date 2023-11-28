package OMarketEventListener

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	pkg "service/pkg/smart-contracts"
	"time"

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

func formatPublicOrderCreatedEvent(event *pkg.OpenMarketPublicOrderCreated) string {
	// Formatação básica da saída do evento
	return fmt.Sprintf(
		"Public Order Created Event:\n"+
			"    Token ID: %d\n"+
			"    Transaction Hash: %s\n"+
			"    Block Number: %d\n"+
			"    Block Hash: %s\n"+
			"    Event Index: %d\n",
		event.TokenId,
		event.Raw.TxHash.Hex(),
		event.Raw.BlockNumber,
		event.Raw.BlockHash.Hex(),
		event.Raw.Index,
	)
}

func (cr *ContractRepository) ListenToPublicOrderCreated(ctx context.Context) {
	go func() {
		for {
			currentBlock, err := cr.GetCurrentBlockNumber()
			if err != nil {
				log.Printf("Error getting current block number: %v", err)
				time.Sleep(time.Second * 10)
				continue
			}

			log.Printf("Current Block: %v", currentBlock)

			opts := &bind.WatchOpts{Start: &currentBlock, Context: ctx}
			events := make(chan *pkg.OpenMarketPublicOrderCreated)
			var tokenIdFilter []*big.Int

			sub, err := cr.contract.WatchPublicOrderCreated(opts, events, tokenIdFilter)
			if err != nil {
				log.Printf("Failed to watch event logs: %v", err)
				time.Sleep(time.Second * 5)
				continue
			}

			for {
				select {
				case event := <-events:
					formattedEvent := formatPublicOrderCreatedEvent(event)
					log.Println(formattedEvent)
				case err := <-sub.Err():
					log.Printf("Subscription error: %v", err)
					break
				case <-ctx.Done():
					return
				}
			}

			sub.Unsubscribe()
			time.Sleep(time.Second * 5)
		}
	}()
}

func formatPrimarySaleEvent(event *pkg.OpenMarketPrimarySale) string {
	return fmt.Sprintf(
		"Primary Sale Event:\n"+
			"\tSender: %s\n"+
			"\tToken ID: %d\n"+
			"\tAmount: %d\n"+
			"\tTransaction Hash: %s\n"+
			"\tBlock Number: %d\n",
		event.Sender.Hex(),     // Converte o endereço para string
		event.TokenId,          // ID do token
		event.Amount,           // Quantidade
		event.Raw.TxHash.Hex(), // Hash da transação
		event.Raw.BlockNumber,  // Número do bloco
	)
}
func (cr *ContractRepository) ListenToPrimarySale(ctx context.Context) {
	go func() {
		for {
			currentBlock, err := cr.GetCurrentBlockNumber()
			if err != nil {
				log.Printf("Error getting current block number: %v", err)
				time.Sleep(time.Second * 10)
				continue
			}

			log.Printf("Current Block: %v", currentBlock)

			// Channel go to receive events
			events := make(chan *pkg.OpenMarketPrimarySale)
			opts := &bind.WatchOpts{Start: &currentBlock, Context: ctx}

			// Filters
			var senderFilter []common.Address
			var tokenIdFilter []*big.Int
			var amountFilter []*big.Int

			// Init Observer
			sub, err := cr.contract.WatchPrimarySale(opts, events, senderFilter, tokenIdFilter, amountFilter)
			if err != nil {
				log.Printf("Failed to start watching event logs: %v", err)
				time.Sleep(time.Second * 5)
				continue
			}

			// Process Events
			for {
				select {
				case event := <-events:
					formattedEvent := formatPrimarySaleEvent(event)
					log.Println(formattedEvent)
				case err := <-sub.Err():
					log.Printf("Subscription error: %v", err)
					break
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				}
			}
		}
	}()
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
