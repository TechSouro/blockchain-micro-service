// ApiModule/api_handler.go
package ApiModule

import (
	"encoding/json"
	"net/http"
	usecase "service/internal/ApiModule/usecase"
)

type ApiHandler struct {
	useCase *usecase.ApiUseCase
}

func NewApiHandler(useCase *usecase.ApiUseCase) *ApiHandler {
	return &ApiHandler{
		useCase: useCase,
	}
}

func (ah *ApiHandler) GetTreasuryEventsHandler(w http.ResponseWriter, r *http.Request) {
	treasuryEvents, err := ah.useCase.GetTreasuryEvents()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(treasuryEvents)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func (ah *ApiHandler) GetPrimaryTableHandler(w http.ResponseWriter, r *http.Request) {
	primaryOrders, err := ah.useCase.GetPrimaryTable()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(primaryOrders)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func (ah *ApiHandler) GetAllSecondaryTableItemsHandler(w http.ResponseWriter, r *http.Request) {
	secondaryOrders, err := ah.useCase.GetAllSecondaryTableItems()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(secondaryOrders)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func (ah *ApiHandler) GetAllTransferTableItemsHandler(w http.ResponseWriter, r *http.Request) {
	transferItems, err := ah.useCase.GetAllTransferTableItems()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(transferItems)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
