package OracleDrex

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	pkg "service/pkg/smart-contracts"
	strings "strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TransferTable struct {
	ID        uint64 `gorm:"primaryKey"`
	TxOrigin  string `gorm:"column:tx_origin"`
	Receiver  string `gorm:"column:receiver"`
	Amount    uint64 `gorm:"column:amount"`
	Timestamp string `gorm:"column:timestamp"`
}

type EventTransfer struct {
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Amount    *big.Int       `json:"amount"`
	TxOrigin  common.Address `json:"tx_origin"`
	Receiver  common.Address `json:"receiver"`
	Timestamp uint64         `json:"timestamp"`
	TxHash    common.Hash    `json:"tx_hash"`
}

type ContractRepository struct {
	contract *pkg.OracleDrexerc3643
	client   *ethclient.Client
	DB       *gorm.DB
}

// Função para formatar o evento Transfer em um objeto JSON
func formatTransferEvent(client *ethclient.Client, event *pkg.OracleDrexerc3643Transfer) ([]byte, error) {
	// Obtenha o cabeçalho do bloco para obter a hora exata da transação
	header, err := client.HeaderByHash(context.Background(), event.Raw.BlockHash)
	if err != nil {
		log.Printf("Error getting block header: %v", err)
		return nil, err
	}

	// Crie um struct EventTransfer
	transferEvent := EventTransfer{
		From:      event.From,
		To:        event.To,
		Amount:    event.Value,
		TxOrigin:  common.BytesToAddress(event.Raw.Topics[1].Bytes()),
		Receiver:  common.BytesToAddress(event.Raw.Topics[2].Bytes()),
		TxHash:    event.Raw.TxHash,
		Timestamp: uint64(header.Time),
	}

	// Converta o struct para JSON formatado
	jsonData, err := json.MarshalIndent(transferEvent, "", "    ")
	if err != nil {
		log.Printf("Error formatting Transfer event to JSON: %v", err)
		return nil, err
	}

	return jsonData, nil
}

func NewContractRepository(clientURL string, contractAddress common.Address, dbURL string) (*ContractRepository, error) {
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the PostgreSQL database: %v", err)
		return nil, err
	}

	// Configura a automigração para criar tabelas automaticamente
	err = gormDB.AutoMigrate(&TransferTable{})
	if err != nil {
		log.Fatalf("Failed to automigrate tables: %v", err)
		return nil, err
	}

	contract, err := pkg.NewOracleDrexerc3643(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to bind to the deployed instance of the contract: %v", err)
		return nil, err
	}

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

// event
func (cr *ContractRepository) ListenTransferEvent(ctx context.Context, targetAddress common.Address) {
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
			events := make(chan *pkg.OracleDrexerc3643Transfer)

			var from []common.Address
			var to []common.Address

			sub, err := cr.contract.WatchTransfer(opts, events, from, to)
			if err != nil {
				log.Printf("Failed to watch event logs: %v", err)
				time.Sleep(time.Second * 5)
				continue
			}

			for {
				select {
				case event := <-events:
					// Verifique se o remetente é igual ao endereço alvo
					if event.From == targetAddress {
						formattedEvent, err := formatTransferEvent(cr.client, event)
						if err != nil {
							log.Printf("Error formatting Transfer event: %v", err)
							continue
						}

						log.Println(string(formattedEvent))

						txOriginReceipt := common.BytesToAddress(event.Raw.Topics[1].Bytes())
						txReceiverReceipt := common.BytesToAddress(event.Raw.Topics[2].Bytes())

						// save in the database
						cr.SaveTransferEvent(ctx, TransferTable{
							TxOrigin:  txOriginReceipt.String(),
							Receiver:  txReceiverReceipt.String(),
							Amount:    event.Value.Uint64(),
							Timestamp: strings.FormatUint(uint64(time.Now().Unix()), 10),
						})

					}
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

func (repo *ContractRepository) SaveTransferEvent(ctx context.Context, item TransferTable) {

	tx := repo.DB.Begin()

	transferOmitted := &TransferTable{
		TxOrigin:  item.TxOrigin,
		Receiver:  item.Receiver,
		Amount:    item.Amount,
		Timestamp: item.Timestamp,
	}

	fmt.Printf("Transfer Event Listed: %v", transferOmitted)

	// Salve os dados na tabela
	result := tx.Create(transferOmitted)

	if result.Error != nil {
		log.Println("Error adding to primary_table:", result.Error)

		tx.Rollback()
		return
	}

	// Commita a transação
	tx.Commit()
}
