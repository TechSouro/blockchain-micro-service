package OracleDrex

import (
	"context"
	"log"
	repository "service/internal/OracleDrex/repository"
	"time"
)

type TransferEventUseCase struct {
	repo *repository.ContractRepository
}

func NewTransferEventUseCase(repo *repository.ContractRepository) *TransferEventUseCase {
	return &TransferEventUseCase{
		repo: repo,
	}
}

func (uc *TransferEventUseCase) ProcessListenTransferEvent() {
	go func() {
		defer uc.recoverFromError("ProcessListenTransferEvent")
		ctx := context.Background()
		uc.repo.ListenTransferEvent(ctx)
	}()
}

func (uc *TransferEventUseCase) recoverFromError(processName string) {
	if r := recover(); r != nil {
		log.Printf("Recovered from error in %s: %v", processName, r)
		time.Sleep(time.Second * 10)
		switch processName {
		case "ProcessListenTransferEvent":
			uc.ProcessListenTransferEvent()
		}
	}
}
