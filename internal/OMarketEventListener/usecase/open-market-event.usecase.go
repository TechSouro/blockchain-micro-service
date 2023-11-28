package OMarketEventListener

import (
	"context"
	"log"
	repository "service/internal/OMarketEventListener/repository"
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

func (uc *EventUseCase) ProcessPublicOrderCreated() {
	go func() {
		defer uc.recoverFromError("ProcessPublicOrderCreated")
		ctx := context.Background()
		uc.repo.ListenToPublicOrderCreated(ctx)
	}()
}

func (uc *EventUseCase) ProcessPrimarySale() {
	go func() {
		defer uc.recoverFromError("ProcessPrimarySale")

		ctx := context.Background()
		uc.repo.ListenToPrimarySale(ctx)
	}()
}

func (uc *EventUseCase) ProcessSecondarySold() {
	go func() {
		defer uc.recoverFromError("ProcessSecondarySold")

		ctx := context.Background()
		uc.repo.ListenSecondarySold(ctx)
	}()
}

func (uc *EventUseCase) ProcessSecondaryForSale() {
	go func() {
		defer uc.recoverFromError("ProcessSecondaryForSale")

		ctx := context.Background()

		uc.repo.ListenSecondaryForSale(ctx)
	}()
}

func (uc *EventUseCase) recoverFromError(processName string) {
	if r := recover(); r != nil {
		log.Printf("Recovered from error in %s: %v", processName, r)
		time.Sleep(time.Second * 10)
		switch processName {
		case "ProcessPublicOrderCreated":
			uc.ProcessPublicOrderCreated()
		case "ProcessPrimarySale":
			uc.ProcessPrimarySale()
		case "ProcessSecondarySold":
			uc.ProcessSecondarySold()
		case "ProcessSecondaryForSale":
			uc.ProcessSecondaryForSale()
		}
	}
}
