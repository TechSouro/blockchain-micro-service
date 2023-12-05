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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PrimaryTable struct {
	TokenID   uint64 `gorm:"column:token_id"`
	Available uint64
	Price     string
}

type SecondaryTable struct {
	TokenID   uint64 `gorm:"column:token_id"`
	Seller    string
	Available uint64
	Price     string
}

type PublicOrderCreated struct {
	TokenID   uint64 `gorm:"column:token_id"`
	Available uint64
	Price     *big.Int
}

type PrimarySale struct {
	TokenID uint64 `gorm:"column:token_id"`
	Amount  uint64
}

type SecondaryForSale struct {
	TokenID uint64 `gorm:"column:token_id"`
	Seller  string
	Units   uint64
	Price   *big.Int
}

type SecondarySold struct {
	TokenID uint64 `gorm:"column:token_id"`
	Seller  string
	Buyer   string
	Units   uint64
	Price   uint64
}

type ContractRepository struct {
	contract *pkg.OpenMarket
	client   *ethclient.Client
	DB       *gorm.DB
}

func NewContractRepository(clientURL string, contractAddress common.Address, dbURL string) (*ContractRepository, error) {
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}
	log.Printf("Ethereum client connected successfully")

	gormDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the PostgreSQL database: %v", err)
		return nil, err
	}

	// Configura a automigração para criar tabelas automaticamente
	err = gormDB.AutoMigrate(&PrimaryTable{}, &SecondaryTable{})
	if err != nil {
		log.Fatalf("Failed to automigrate tables: %v", err)
		return nil, err
	}

	contract, err := pkg.NewOpenMarket(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to bind to the deployed instance of the contract: %v", err)
		return nil, err
	}

	log.Printf("PostgreSQL database connected successfully")

	return &ContractRepository{contract: contract, client: client, DB: gormDB}, nil
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
		yellow("Price: "), formatPriceBigInt(event.Price),
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
					tokenID := bigIntToUint64(event.TokenId)
					unitsAvailable := bigIntToUint64(event.Units)
					price := event.Price

					cr.AddToPrimaryTable(ctx, PublicOrderCreated{
						TokenID:   tokenID,
						Available: unitsAvailable,
						Price:     price,
					})
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

					tokenID := bigIntToUint64(event.TokenId)

					cr.SubtractFromPrimaryTable(ctx, PrimarySale{
						TokenID: tokenID,
						Amount:  event.Amount.Uint64(),
					})
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
		yellow("Price:"), formatPriceBigInt(event.Price),
		yellow("Block Number:"), event.Raw.BlockNumber,
	)
}

func (cr *ContractRepository) ListenSecondaryForSale(ctx context.Context) {
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
			events := make(chan *pkg.OpenMarketSecondaryForSale)
			opts := &bind.WatchOpts{Start: &currentBlock, Context: ctx}

			// Filters
			var sellerRule []common.Address

			var tokenIdRule []*big.Int

			var unitsRule []*big.Int

			// Init Observer
			sub, err := cr.contract.WatchSecondaryForSale(opts, events, sellerRule, tokenIdRule, unitsRule)
			if err != nil {
				log.Printf("Failed to start watching event logs: %v", err)
				time.Sleep(time.Second * 5)
				continue
			}

			// Process Events
			for {
				select {
				case event := <-events:
					formattedEvent := formatSecondaryForSaleEvent(event)
					log.Println(formattedEvent)

					tokenID := bigIntToUint64(event.TokenId)

					cr.AddToSecondaryTable(ctx, SecondaryForSale{
						TokenID: tokenID,
						Seller:  event.Seller.Hex(),
						Units:   event.Units.Uint64(),
						Price:   event.Price,
					})
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
		yellow("Price:"), formatPriceBigInt(event.Price),
		yellow("Token ID:"), event.TokenId,
	)
}

func formatBigInt(value *big.Int) string {
	if value == nil {
		return "0"
	}
	return value.String()
}

func bigIntToUint64(value *big.Int) uint64 {
	if value == nil {
		return 0
	}
	return value.Uint64()
}

func (cr *ContractRepository) ListenSecondarySold(ctx context.Context) {
	go func() {
		for {
			fmt.Printf("Listening for secundary sold events...\n")
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

					tokenID := bigIntToUint64(event.TokenId)

					cr.SubtractFromSecondaryTable(ctx, SecondarySold{
						TokenID: tokenID,
						Seller:  event.Seller.Hex(),
						Buyer:   event.Buyer.Hex(),
						Units:   event.Units.Uint64(),
						Price:   event.Price.Uint64(),
					})
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

/** DATABASE /*
 *
 *
 */
func (repo *ContractRepository) AddToPrimaryTable(ctx context.Context, item PublicOrderCreated) {
	priceValueDTO := formatPriceBigInt(item.Price)
	fmt.Printf("DTO Price value: %v", item.Price)

	primaryOrder := &PrimaryTable{
		TokenID:   item.TokenID,
		Available: item.Available,
		Price:     priceValueDTO,
	}

	fmt.Printf("Primary order: %v", primaryOrder)
	result := repo.DB.Create(primaryOrder)
	if result.Error != nil {
		log.Println("Error adding to primary_table:", result.Error)
	}
}

func (repo *ContractRepository) SubtractFromPrimaryTable(ctx context.Context, item PrimarySale) {
	result := repo.DB.Model(&PrimaryTable{}).
		Where("token_id = ?", item.TokenID).
		Update("Available", gorm.Expr("Available - ?", item.Amount))
	if result.Error != nil {
		log.Println("Error subtracting from primary_table:", result.Error)
	}
}

func (repo *ContractRepository) AddToSecondaryTable(ctx context.Context, item SecondaryForSale) {
	secondaryOrder := &SecondaryTable{
		TokenID:   item.TokenID,
		Seller:    item.Seller,
		Available: item.Units,
		Price:     formatPriceBigInt(item.Price),
	}
	result := repo.DB.Create(secondaryOrder)
	if result.Error != nil {
		log.Println("Error adding to secondary_table:", result.Error)
	}
}

func (repo *ContractRepository) SubtractFromSecondaryTable(ctx context.Context, item SecondarySold) {
	result := repo.DB.Model(&SecondaryTable{}).
		Where("token_id = ?", item.TokenID).
		Update("Available", gorm.Expr("Available - ?", item.Units))
	if result.Error != nil {
		log.Println("Error subtracting from secondary_table:", result.Error)
	}
}

// func formatPriceBigInt(price *big.Int) string {

// 	price.Div(price, new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil))

// 	priceFloat := new(big.Float).SetInt(price)
// 	priceFloat.Quo(priceFloat, big.NewFloat(100))

// 	priceStr := fmt.Sprintf("%.2f", priceFloat)

//		return priceStr
//	}
func formatPriceBigInt(price *big.Int) string {
	// Divide o preço por 10^16
	price.Div(price, new(big.Int).Exp(big.NewInt(10), big.NewInt(16), nil))

	// Cria um big.Float a partir do preço
	priceFloat := new(big.Float).SetInt(price)

	// Divide o preço float por 100 (para converter para porcentagem)
	priceFloat.Quo(priceFloat, big.NewFloat(100))

	// Converte o preço float para uma string formatada
	return fmt.Sprintf("%.2f", priceFloat)
}
