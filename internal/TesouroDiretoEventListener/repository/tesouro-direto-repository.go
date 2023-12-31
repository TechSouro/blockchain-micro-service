package OMarketEventListener

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"math/big"
	pkg "service/pkg/smart-contracts"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ContractRepository struct {
	contract *pkg.TesouroDireto
	client   *ethclient.Client
	DB       *gorm.DB
}

type Bigint struct {
	*big.Int
}

type TreasuryCreatedTable struct {
	TokenID      uint64
	TotalValue   Bigint
	APY          uint64
	Duration     uint64
	TreasuryType uint64
}

// Value implements the driver.Valuer interface
func (i Bigint) Value() (driver.Value, error) {
	return i.Int.String(), nil
}

func (i *Bigint) Scan(src interface{}) error {
	switch src := src.(type) {
	case int64:
		i.Int = big.NewInt(src)
	case []byte:
		i.Int.SetString(string(src), 10)
	default:
		return driver.ErrSkip
	}
	return nil
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
	err = gormDB.AutoMigrate(&TreasuryCreatedTable{})
	if err != nil {
		log.Fatalf("Failed to automigrate tables: %v", err)
		return nil, err
	}

	contract, err := pkg.NewTesouroDireto(contractAddress, client)
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

func (cr *ContractRepository) ListenEmitTresureryCreated(ctx context.Context) {
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
			events := make(chan *pkg.TesouroDiretoTreasuryCreated)

			var tokenIdRule []*big.Int
			var totalValueRule []*big.Int
			var apyRule []*big.Int

			sub, err := cr.contract.WatchTreasuryCreated(opts, events, tokenIdRule, totalValueRule, apyRule)
			if err != nil {
				log.Printf("Failed to watch event logs: %v", err)
				time.Sleep(time.Second * 5)
				continue
			}

			for {
				select {
				case event := <-events:
					tokenID := bigIntToUint64(event.TokenId)
					apy := bigIntToUint64(event.Apy)
					duration := bigIntToUint64(event.Duration)
					treasuryType := event.Arg4

					// Crie uma instância de Bigint para TotalValue
					totalValue := Bigint{Int: event.TotalValue}

					cr.saveTreasuryCreated(ctx, TreasuryCreatedTable{
						TokenID:      tokenID,
						TotalValue:   totalValue,
						APY:          apy,
						Duration:     duration,
						TreasuryType: uint64(treasuryType),
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

/** DATABASE /*
 *
 *
 */
func (repo *ContractRepository) saveTreasuryCreated(ctx context.Context, item TreasuryCreatedTable) {

	treasureEvent := &TreasuryCreatedTable{
		TokenID:      item.TokenID,
		TotalValue:   item.TotalValue,
		APY:          item.APY,
		Duration:     item.Duration,
		TreasuryType: item.TreasuryType,
	}

	fmt.Printf("Treasure created event: %v", treasureEvent)
	result := repo.DB.Create(treasureEvent)
	if result.Error != nil {
		log.Println("Error adding to primary_table:", result.Error)
	}
}

func bigIntToUint64(value *big.Int) uint64 {
	if value == nil {
		return 0
	}
	return value.Uint64()
}
