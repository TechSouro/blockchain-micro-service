package OracleDrex

import (
	"context"
	"log"
	repository "service/internal/OracleDrex/repository"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type TransferEventUseCase struct {
	repo *repository.ContractRepository
}

func NewTransferEventUseCase(repo *repository.ContractRepository) *TransferEventUseCase {
	return &TransferEventUseCase{
		repo: repo,
	}
}

func (uc *TransferEventUseCase) ProcessListenTransferEvent(targetAddress common.Address) {
	go func() {
		defer uc.recoverFromError("ProcessListenTransferEvent", targetAddress)
		ctx := context.Background()
		uc.repo.ListenTransferEvent(ctx, targetAddress)
	}()
}

func (uc *TransferEventUseCase) recoverFromError(processName string, targetAddress common.Address) {
	if r := recover(); r != nil {
		log.Printf("Recovered from error in %s: %v", processName, r)
		time.Sleep(time.Second * 10)
		switch processName {
		case "ProcessListenTransferEvent":
			uc.ProcessListenTransferEvent(targetAddress)
		}
	}
}
