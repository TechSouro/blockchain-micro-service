package config

import (
	"fmt"
	"log"
	"net/http"
	OMarketEventListenerRepository "service/internal/OMarketEventListener/repository"
	OMarketEventListenerUseCase "service/internal/OMarketEventListener/usecase"
	OracleDrexRepository "service/internal/OracleDrex/repository"
	OracleDrexUseCase "service/internal/OracleDrex/usecase"
	TesouroDiretoRepository "service/internal/TesouroDiretoEventListener/repository"
	TesouroDiretoUseCase "service/internal/TesouroDiretoEventListener/usecase"

	"github.com/ethereum/go-ethereum/common"
)

func InitEventSystem() {
	// Ethereum client

	// clientURL := "wss://sepolia.infura.io/ws/v3/60786ed4ffd74c75b4b0bb369cde55f7"
	clientURL := "ws://127.0.01:6175"
	contractAddr := common.HexToAddress("0x232efcE3E92a8126Ae6A91ECb78844b8D607Ed71")
	drexContractAddr := common.HexToAddress("0xF8fFf65FA486769ff4486d6B79a7D1f213E0Ff9C")
	vaultAddressToTarget := common.HexToAddress("0x4Eb845fc5eedcf3f1a7925F47372a3a9aa437adE")
	tesouroDiretoAddr := common.HexToAddress("0xe0cA2EeE246a72e86F2949F5F01bAb733a09209c")

	// PostgreSQL database URL
	dbURL := "postgres://admin:123@lunave@localhost:5432/blockservice"

	// Instanciate repository OpenMarket Contract
	openMarketRepository, err := OMarketEventListenerRepository.NewContractRepository(clientURL, contractAddr, dbURL)
	if err != nil {
		log.Fatalf("Failed to create contract OpenMarket repository: %v", err)
	}
	// Instanciate repository OpenMarket Contract
	oracleDrexRepository, err := OracleDrexRepository.NewContractRepository(clientURL, drexContractAddr, dbURL)
	if err != nil {
		log.Fatalf("Failed to create contract Drex repository: %v", err)
	}

	// Instanciate Tesouro direto
	tesouroDiretoRepository, err := TesouroDiretoRepository.NewContractRepository(clientURL, tesouroDiretoAddr, dbURL)
	if err != nil {
		log.Fatalf("Failed to create contract Tesouro Direto repository: %v", err)
	}

	openMarketEventUseCase := OMarketEventListenerUseCase.NewEventUseCase(openMarketRepository)
	oracleDrexEventUseCase := OracleDrexUseCase.NewTransferEventUseCase(oracleDrexRepository)
	tesouroDiretoEventUsecase := TesouroDiretoUseCase.NewEventUseCase(tesouroDiretoRepository)

	// Initialize events
	go openMarketEventUseCase.ProcessPublicOrderCreated()
	go openMarketEventUseCase.ProcessPrimarySale()
	go openMarketEventUseCase.ProcessSecondaryForSale()
	go openMarketEventUseCase.ProcessSecondarySold()

	// initialize oracleDrex events
	go oracleDrexEventUseCase.ProcessListenTransferEvent(vaultAddressToTarget)
	go tesouroDiretoEventUsecase.ProcessListenEmitTresure()

	// Inicializa as rotas da API
	r := InitializeAPI(dbURL)

	// Inicia o servidor
	port := 4000
	fmt.Printf("Server listening on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)

	log.Println("Event system initialized and listening for events")

}
