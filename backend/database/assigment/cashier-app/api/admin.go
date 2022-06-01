package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ruang-guru/playground/backend/database/assigment/cashier-app/repository"
)

type AdminErrorResponse struct {
	Error string `json:"error"`
}

type ProductSales struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Quantity int    `json:"quantity"`
}

type AdminResponse struct {
	Sales []repository.Sales `json:"sales"`
}

func (api *API) getDashboard(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	productName := req.URL.Query().Get("product_name")

	startPeriod, _ := time.Parse("2006-01-02", req.URL.Query().Get("start_period"))
	endPeriod, _ := time.Parse("2006-01-02", req.URL.Query().Get("end_period"))

	getSalesRequest := repository.GetSalesRequest{
		ProductName: productName,
		StartPeriod: &startPeriod,
		EndPeriod:   &endPeriod,
	}

	if req.URL.Query().Get("start_period") == "" {
		getSalesRequest.StartPeriod = nil
	}
	if req.URL.Query().Get("end_period") == "" {
		getSalesRequest.EndPeriod = nil
	}

	encoder := json.NewEncoder(w)

	sales, err := api.salesRepo.FetchSales(getSalesRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(CartErrorResponse{Error: err.Error()})
		return
	}

	encoder.Encode(AdminResponse{Sales: sales})
	w.WriteHeader(http.StatusOK)
}
