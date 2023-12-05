// ApiModule/api_routes.go
package ApiModule

import (
	handler "service/internal/ApiModule"

	"github.com/gorilla/mux"
)

func ConfigureAPIRoutes(apiHandler *handler.ApiHandler) *mux.Router {
	r := mux.NewRouter()

	// Endpoint para listar os eventos do tesouro
	r.HandleFunc("/api/treasurys", apiHandler.GetTreasuryEventsHandler).Methods("GET")
	r.HandleFunc("/api/primarysale", apiHandler.GetPrimaryTableHandler).Methods("GET")
	r.HandleFunc("/api/secondarysale", apiHandler.GetAllSecondaryTableItemsHandler).Methods("GET")

	return r
}
