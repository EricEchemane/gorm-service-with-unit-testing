package main

import (
	"fmt"
	"gopher/infra/db/dbimpl"
	"gopher/services/product"
)

func main() {
	db := dbimpl.New(&product.Product{})
	service := product.NewService(db)
	products, _ := service.GetProducts(100)

	for _, p := range products {
		fmt.Println("Product code:", p.Code)
		fmt.Println("Product price:", p.Price)
		fmt.Println("-------------------------")
	}
}
