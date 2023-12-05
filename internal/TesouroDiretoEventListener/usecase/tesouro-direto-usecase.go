package OMarketEventListener

import (
	"context"
	"log"
	repository "service/internal/TesouroDiretoEventListener/repository"
	"time"
)

type EventUseCase struct {
	repo *repository.ContractRepository
}

func NewEventUseCase(repo *repository.ContractRepository) *EventUseCase {
	return &EventUseCase{
		repo: repo,
	}
}

func (uc *EventUseCase) ProcessListenEmitTresure() {
	go func() {
		defer uc.recoverFromError("ProcessListenEmitTresure")
		ctx := context.Background()
		uc.repo.ListenEmitTresureryCreated(ctx)
	}()
}

func (uc *EventUseCase) recoverFromError(processName string) {
	if r := recover(); r != nil {
		log.Printf("Recovered from error in %s: %v", processName, r)
		time.Sleep(time.Second * 10)
		switch processName {
		case "ProcessListenEmitTresure":
			uc.ProcessListenEmitTresure()
		}
	}
}
