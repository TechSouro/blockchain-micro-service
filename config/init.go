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
	contractAddr := common.HexToAddress("0xe50f22f357DCCA5b1950FD1C7F1A595d42BC696C")

	// Instanciate repository
	repo, err := OMarketEventListenerRepository.NewContractRepository(clientURL, contractAddr)
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
