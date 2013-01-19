package model

import (
	"strings"
	"strconv"
	"api"
	// "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	// "fmt"
)

var productTN = "products"

// TODO try to refactor this method
func (model Model) PutProduct(id string, input api.ProductInput) (product *api.Product, err error) {
	// TODO check whether all the input field is nil in order to prevent removing data unintendedly
	// TODO find a better way to assign newProduct map
	newProduct := make(map[string]interface{})
	newProduct["Name"] = input.Name
	newProduct["PhotoLink"] = input.PhotoLink
	newProduct["Price"] = input.Price
	newProduct["Id"] = id
	
	err = model.put(productTN, newProduct, "Id", genId())
	if err != nil {
		return
	}
	
	// TODO find out why I have to use map to unmarshal data if I insert it with a map
	product = &api.Product{}
	productMap := map[string]interface{}{}
	err = db.C(productTN).Find(bson.M{"Id": id}).One(productMap)
	if err != nil {
		return
	}
	product.Name = productMap["Name"].(string)
	product.PhotoLink = productMap["PhotoLink"].(string)
	product.Price = productMap["Price"].(int64)
	product.Id = productMap["Id"].(string)
	
	return
}

// TODO remove this function and try to use bson.ObjectId 
func genId() string {
	productC := db.C(productTN)
	count, err := productC.Count()
	if err != nil {
		panic(err)
	}

	no := []string{"no", strconv.Itoa(count)}
	return strings.Join(no, "_")
}

func (model Model) RemoveProduct(id string) (err error) {
	err = model.remove(productTN, "Id", id)
	return
}