package main

import (
	"model"
	"api"
)

func main() {
	mymodel := model.Model{}
	productInput := api.ProductInput{"empty...", 30, "photolink"}
	product, _ := mymodel.PutProduct("no_18", productInput)
	print(product.Id, "\n")
	print(product.Id, "\n")
	// mymodel.RemoveProduct("no_13")
}