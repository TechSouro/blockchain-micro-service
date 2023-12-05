// ApiModule/usecase/api_usecase.go
package ApiModule

import (
	domain "service/internal/ApiModule/domain"
	repository "service/internal/ApiModule/repository"
)

type ApiUseCase struct {
	repo *repository.ApiRepository
}

func NewApiUseCase(repo *repository.ApiRepository) *ApiUseCase {
	return &ApiUseCase{
		repo: repo,
	}
}

func (uc *ApiUseCase) GetTreasuryEvents() ([]domain.GetTreasuryResponse, error) {
	treasuryEvents, err := uc.repo.GetTreasuryEvents()
	if err != nil {
		return nil, err
	}

	return treasuryEvents, nil
}

func (uc *ApiUseCase) GetPrimaryTable() ([]domain.PrimaryCreatedTable, error) {
	return uc.repo.GetAllPrimaryTableItems()
}

func (uc *ApiUseCase) GetAllSecondaryTableItems() ([]domain.SecondaryCreatedTable, error) {
	secondaryOrders, err := uc.repo.GetAllSecondaryTableItems()
	if err != nil {
		return nil, err
	}

	return secondaryOrders, nil
}

func (uc *ApiUseCase) GetAllTransferTableItems() ([]domain.TransferCreatedTable, error) {
	transferItems, err := uc.repo.GetAllTransferTableItems()
	if err != nil {
		return nil, err
	}

	return transferItems, nil
}
