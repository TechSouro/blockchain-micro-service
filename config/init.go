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
	clientURL := "wss://sepolia.infura.io/ws/v3/60786ed4ffd74c75b4b0bb369cde55f7"
	contractAddr := common.HexToAddress("0x23f455bCFFA666cBE80e40bF81B8A32e8E6D6759")
	drexContractAddr := common.HexToAddress("0x898c490d2AabCD87BC3E934cB708482A7532b887")
	vaultAddressToTarget := common.HexToAddress("0xD6ecc456121f1a72a3d01a3144Fb74bB31fAF571")

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
