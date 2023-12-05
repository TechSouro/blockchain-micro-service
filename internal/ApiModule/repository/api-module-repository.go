// ApiModule/repository/api_repository.go
package ApiModule

import (
	domain "service/internal/ApiModule/domain"

	"gorm.io/gorm"
)

type ApiRepository struct {
	DB *gorm.DB
}

func NewApiRepository(db *gorm.DB) *ApiRepository {
	return &ApiRepository{
		DB: db,
	}
}

func (ar *ApiRepository) GetTreasuryEvents() ([]domain.GetTreasuryResponse, error) {
	var results []domain.TreasuryCreatedTable
	if err := ar.DB.Table("treasury_created_tables").Find(&results).Error; err != nil {
		return nil, err
	}

	var apiResponseList []domain.GetTreasuryResponse
	for _, result := range results {
		apiResponseList = append(apiResponseList, TreasuryCreatedTableToApiResponse(result))
	}

	return apiResponseList, nil
}

func TreasuryCreatedTableToApiResponse(treasury domain.TreasuryCreatedTable) domain.GetTreasuryResponse {
	return domain.GetTreasuryResponse{
		TokenID:      treasury.TokenID,
		TotalValue:   treasury.TotalValue,
		APY:          treasury.APY,
		Duration:     treasury.Duration,
		TreasuryType: treasury.TreasuryType,
	}
}

func (repo *ApiRepository) GetAllPrimaryTableItems() ([]domain.PrimaryCreatedTable, error) {
	var primaryOrders []domain.PrimaryCreatedTable
	if err := repo.DB.Table("primary_tables").Find(&primaryOrders).Error; err != nil {
		return nil, err
	}

	return primaryOrders, nil
}

func (repo *ApiRepository) GetAllSecondaryTableItems() ([]domain.SecondaryCreatedTable, error) {
	var secondaryOrders []domain.SecondaryCreatedTable
	if err := repo.DB.Table("secondary_tables").Find(&secondaryOrders).Error; err != nil {
		return nil, err
	}

	return secondaryOrders, nil
}
