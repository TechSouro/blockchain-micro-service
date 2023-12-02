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
	"github.com/fatih/color"
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

// ------------------------- Public Order Created Event ----------------------------
func formatPublicOrderCreatedEvent(event *pkg.OpenMarketPublicOrderCreated) string {
	black := color.New(color.BgHiBlack).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	return fmt.Sprintf(
		"%s\n\t %s %d\n %s %d\n %s %s\n \t%s %s\n \t%s %d\n \t%s %s\n \t%s %d\n",
		black("Public Order Created Event:"),
		yellow("Token ID: "), event.TokenId,
		yellow("Units: "), event.Units,
		yellow("Price: "), formatBigInt(event.Price),
		yellow("Transaction Hash:"), event.Raw.TxHash.Hex(),
		yellow("Block Number:"), event.Raw.BlockNumber,
		yellow("Block Hash:"), event.Raw.BlockHash.Hex(),
		yellow("Event Index:"), event.Raw.Index,
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

			var tokenId []*big.Int
			var units []*big.Int

			sub, err := cr.contract.WatchPublicOrderCreated(opts, events, tokenId, units)
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

// ------------------------- primaryForSale Event ----------------------------
func formatPrimarySaleEvent(event *pkg.OpenMarketPrimarySale) string {
	cyan := color.New(color.BgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	return fmt.Sprintf(
		"%s\n\t%s %s\n\t%s %d\n\t%s %d\n\t%s %s\n\t%s %d\n",
		cyan("Primary Sale Event:"),
		yellow("Sender:"), event.Sender.Hex(),
		yellow("Token ID:"), event.TokenId,
		yellow("Amount:"), event.Amount,
		yellow("Transaction Hash:"), event.Raw.TxHash.Hex(),
		yellow("Block Number:"), event.Raw.BlockNumber,
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

// ------------------------- secundaryForSale Event ----------------------------

func formatSecondaryForSaleEvent(event *pkg.OpenMarketSecondaryForSale) string {
	green := color.New(color.BgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	return fmt.Sprintf(
		"%s\n\t%s %s\n\t%s %d\n\t%s %d\n\t%s %s\n\t%s %d\n",
		green("Secondary For Sale Event:"),
		yellow("Seller:"), event.Seller.Hex(),
		yellow("Token ID:"), event.TokenId,
		yellow("Units:"), event.Units,
		yellow("Price:"), formatBigInt(event.Price),
		yellow("Block Number:"), event.Raw.BlockNumber,
	)
}

func (cr *ContractRepository) ListenSecondaryForSale(ctx context.Context) {
	go func() {
		for {
			currentBlock, err := cr.GetCurrentBlockNumber()
			if err != nil {
				log.Printf("Error getting current block number to secundary sale: %v", err)
				time.Sleep(time.Second * 10)
				continue
			}

			// Channel go to receive events
			events := make(chan *pkg.OpenMarketSecondaryForSale)
			opts := &bind.WatchOpts{Start: &currentBlock, Context: ctx}

			// Filters
			var sellersFilter []common.Address
			var tokenIdsFilter []*big.Int
			var amountsFilter []*big.Int

			// Init Observer
			sub, err := cr.contract.WatchSecondaryForSale(opts, events, sellersFilter, tokenIdsFilter, amountsFilter)
			if err != nil {
				log.Printf("Failed to start watching event logs for secundary Sale: %v", err)
				time.Sleep(time.Second * 5)
				continue
			}

			// Process Events
			for {
				select {
				case event := <-events:
					formattedEvent := formatSecondaryForSaleEvent(event)
					fmt.Println(formattedEvent)
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

func formatSecondarySoldEvent(event *pkg.OpenMarketSecondarySold) string {
	magent := color.New(color.BgMagenta).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	return fmt.Sprintf(
		"%s\n\t%s %s\n\t%s %s\n\t%s %d\n\t%s %s\n\t%s %d\n",
		magent("Secondary Sold Event:"),
		yellow("Seller:"), event.Seller.Hex(),
		yellow("Buyer:"), event.Buyer.Hex(),
		yellow("Units:"), event.Units,
		yellow("Price:"), formatBigInt(event.Price),
		yellow("Token ID:"), event.TokenId,
	)
}

func formatBigInt(value *big.Int) string {
	if value == nil {
		return "0"
	}
	return value.String()
}

func (cr *ContractRepository) ListenSecondarySold(ctx context.Context) {
	go func() {
		for {
			currentBlock, err := cr.GetCurrentBlockNumber()
			if err != nil {
				log.Printf("Error getting current block number to secundary sold: %v", err)
				time.Sleep(time.Second * 10)
				continue
			}

			// Channel go to receive events
			events := make(chan *pkg.OpenMarketSecondarySold)
			opts := &bind.WatchOpts{Start: &currentBlock, Context: ctx}

			// Filters
			var sellerRule []common.Address
			var buyerRule []common.Address
			var unitsRule []*big.Int

			// Init Observer
			sub, err := cr.contract.WatchSecondarySold(opts, events, sellerRule, buyerRule, unitsRule)
			if err != nil {
				fmt.Printf("Failed to start watching event logs for secundary sold: %v", err)
				time.Sleep(time.Second * 5)
				continue
			}

			// Process Events
			for {
				select {
				case event := <-events:
					formattedEvent := formatSecondarySoldEvent(event)
					log.Println(formattedEvent)
				case err := <-sub.Err():
					fmt.Printf("Subscription error: %v", err)
					break
				case <-ctx.Done():
					sub.Unsubscribe()
					return
				}
			}
		}
	}()
}
