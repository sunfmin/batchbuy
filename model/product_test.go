package model

import (
	"testing"
	"time"
	"fmt"
	// "labix.org/v2/mgo/bson"
)

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
	
	productList, _ := ProductListOfDate(today)
	
	if len(productList) != 2 {
		fmt.Printf("Product List: %s\n", productList)
		t.Error("Can Get Product List Properly")
	}
	
	productCol.RemoveAll(M{})
}