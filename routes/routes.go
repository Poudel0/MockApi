package routes

import (
	"net/http"

	"github.com/MockApis/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}).Methods(http.MethodGet)

	r.HandleFunc("/sila_transact", services.HandleSilaTransact).Methods(http.MethodPost)
	r.HandleFunc("/nchl_transact", services.HandleNCHLTransact).Methods(http.MethodPost)
	r.HandleFunc("/check_fund_status", services.HandleGenericPost).Methods(http.MethodPost)
	r.HandleFunc("/wallet_payout", services.HandleGenericPost).Methods(http.MethodPost)

	return r
}
