package model

import (
	"fmt"
	"testing"
	"time"
	// "labix.org/v2/mgo/bson"
)

// var db = session.DB("low_tea_at_the_plant_test")

func TestCreateProduct(t *testing.T) {
	productInput := ProductInput{"Name", "Photo Link", 100, time.Now().AddDate(0, 0, -10), time.Now().AddDate(0, 0, 10)}
	product := Product{}
	product.Put("non-exited product", productInput)

	result := Product{}
	productCol.FindId(product.Id).One(&result)

	if product != result {
		t.Error("Can't Create New Prodcut")
	}

	productCol.RemoveAll(M{})
}

func TestUpdateProduct(t *testing.T) {
	productCol.RemoveAll(M{})
	productInput := ProductInput{"Name", "Photo Link", 100, time.Now().AddDate(0, 0, -10), time.Now().AddDate(0, 0, 10)}
	product := Product{}
	product.Put("non-exited product", productInput)

	productInput.Name = "new name"
	product.Put(product.Id.Hex(), productInput)

	result := Product{}
	productCol.FindId(product.Id).One(&result)

	if product != result && result.Name != "new name" {
		t.Error("Can't Create New Prodcut")
	}

	productCol.RemoveAll(M{})
}

func TestProductListOfDate(t *testing.T) {
	today := time.Now()
	initDbForProductTests()

	productList, _ := ProductListOfDate(today)

	if len(productList) != 2 {
		fmt.Printf("Product List: %s\n", productList)
		t.Error("Can Get Product List Properly")
	}

	productCol.RemoveAll(M{})
}

func initDbForProductTests() {
	productCol.RemoveAll(M{})
	today := time.Now()
	productInput := ProductInput{"Product1", "Photo Link", 100, today.AddDate(0, 0, -10), today.AddDate(0, 0, 10)}
	product := Product{}
	product.Put("non-exited product", productInput)
	productInput.Name = "Product2"
	productInput.ValidFrom = today.AddDate(0, 0, -5)
	productInput.ValidTo = today
	product.Put("non-exited product", productInput)
	productInput.Name = "Product3"
	productInput.ValidFrom = today.AddDate(0, 1, 0)
	productInput.ValidTo = today.AddDate(0, 1, 10)
	product.Put("non-exited product", productInput)
}

func TestAllProdutsForApi(t *testing.T) {
	initDbForProductTests()

	products, _ := AllProductsForApi()

	if len(products) != 3 {
		fmt.Printf("Got: %s\n", products)
		t.Error("Can't get all products properly")
	}

	productCol.RemoveAll(M{})
}
