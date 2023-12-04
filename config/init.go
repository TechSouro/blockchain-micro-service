package config

import (
	"log"
	OMarketEventListenerRepository "service/internal/OMarketEventListener/repository"
	OMarketEventListenerUseCase "service/internal/OMarketEventListener/usecase"
	OracleDrexRepository "service/internal/OracleDrex/repository"
	OracleDrexUseCase "service/internal/OracleDrex/usecase"

	"github.com/ethereum/go-ethereum/common"
)

func InitEventSystem() {
	// Ethereum client
	clientURL := "ws://127.0.01:6175"
	contractAddr := common.HexToAddress("0xAdDBA9A148680AA3d00dDcF31CFa3Ec9226de100")
	drexContractAddr := common.HexToAddress("0x10D6995F33Aa5A8e0cCbb882eBE30B51950FC118")
	vaultAddressToTarget := common.HexToAddress("0x809AB5997Ca2D69470f3a154d653ab096d818077")

	// PostgreSQL database URL
	dbURL := "postgres://admin:123@lunave@localhost:5432/blockservice"

	// Instanciate repository OpenMarket Contract
	openMarketRepository, err := OMarketEventListenerRepository.NewContractRepository(clientURL, contractAddr, dbURL)
	if err != nil {
		log.Fatalf("Failed to create contract repository: %v", err)
	}
	// Instanciate repository OpenMarket Contract
	oracleDrexRepository, err := OracleDrexRepository.NewContractRepository(clientURL, drexContractAddr, dbURL)
	if err != nil {
		log.Fatalf("Failed to create contract repository: %v", err)
	}

	openMarketEventUseCase := OMarketEventListenerUseCase.NewEventUseCase(openMarketRepository)
	oracleDrexEventUseCase := OracleDrexUseCase.NewTransferEventUseCase(oracleDrexRepository)

	// Initialize events
	go openMarketEventUseCase.ProcessPublicOrderCreated()
	go openMarketEventUseCase.ProcessPrimarySale()
	go openMarketEventUseCase.ProcessSecondaryForSale()
	go openMarketEventUseCase.ProcessSecondarySold()

	// initialize oracleDrex events
	go oracleDrexEventUseCase.ProcessListenTransferEvent(vaultAddressToTarget)

	log.Println("Event system initialized and listening for events")
}
