package api

import (
	"encoding/json"
	"net/http"
)

type ProductListErrorResponse struct {
	Error string `json:"error"`
}

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

type ProductListSuccessResponse struct {
	Products []Product `json:"products"`
}

func (api *API) productList(w http.ResponseWriter, req *http.Request) {
	productList, err := api.productsRepo.SelectAll()
	encoder := json.NewEncoder(w)

	if err != nil {
		encoder.Encode(CartErrorResponse{err.Error()})
		return
	}
	productResp := make([]Product, 0)
	for _, val := range productList {
		productResp = append(productResp, Product{
			Name:     val.ProductName,
			Category: val.Category,
			Price:    val.Price,
		})
	}
	encoder.Encode(ProductListSuccessResponse{Products: productResp}) // TODO: replace this
}
