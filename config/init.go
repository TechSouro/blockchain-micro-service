package config

import (
	"log"
	OMarketEventListenerRepository "service/internal/OMarketEventListener/repository"
	OMarketEventListenerUseCase "service/internal/OMarketEventListener/usecase"

	"github.com/ethereum/go-ethereum/common"
)

func InitEventSystem() {
	// Ethereum client
	clientURL := "ws://127.0.01:6175"
	contractAddr := common.HexToAddress("0xAdDBA9A148680AA3d00dDcF31CFa3Ec9226de100")

	// PostgreSQL database URL
	dbURL := "postgres://admin:123@lunave@localhost:5432/blockservice"

	// Instanciate repository
	repo, err := OMarketEventListenerRepository.NewContractRepository(clientURL, contractAddr, dbURL)
	if err != nil {
		log.Fatalf("Failed to create contract repository: %v", err)
	}

	eventUseCase := OMarketEventListenerUseCase.NewEventUseCase(repo)

	// Initialize events
	go eventUseCase.ProcessPublicOrderCreated()
	go eventUseCase.ProcessPrimarySale()
	go eventUseCase.ProcessSecondaryForSale()
	go eventUseCase.ProcessSecondarySold()

	log.Println("Event system initialized and listening for events")
}
