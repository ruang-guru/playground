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
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)
	products, err := api.productsRepo.SelectAll()

	if err != nil {
		encoder.Encode(CartErrorResponse{err.Error()})
		return
	}

	productsResp := make([]Product, 0)
	for _, val := range products {
		productsResp = append(productsResp, Product{
			Name:     val.ProductName,
			Price:    val.Price,
			Category: val.Category,
		})
	}

	// fmt.Println(products)

	encoder.Encode(ProductListSuccessResponse{Products: productsResp}) // TODO: replace this
}
