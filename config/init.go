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
	clientURL := "wss://polygon-mumbai.infura.io/ws/v3/60786ed4ffd74c75b4b0bb369cde55f7"
	contractAddr := common.HexToAddress("0x1b77035396c89197509511e9865318329F48B9C3")
	drexContractAddr := common.HexToAddress("0xa56972E78DD14CD54A81A08BE5CA3B43724a07F1")
	vaultAddressToTarget := common.HexToAddress("0x10c45F21464B9EdDb2a82dea04F1D60bba3f73aa")

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
