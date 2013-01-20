package model

import (
	"api"
	"strconv"
	"strings"
	// "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	// "fmt"
)

var productTN = "products"

// TODO try to refactor this method
func (model Model) PutProduct(id string, input api.ProductInput) (product *api.Product, err error) {
	// TODO check whether all the input field is nil in order to prevent removing data unintendedly
	// TODO find a better way to assign newProduct map
	newP := make(map[string]interface{})
	newP["Name"] = input.Name
	newP["PhotoLink"] = input.PhotoLink
	newP["Price"] = input.Price
	newP["Id"] = id

	err = model.put(productTN, newP, "Id", genProductId())
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
	product = newProduct(productMap)

	return
}

// TODO remove this function and try to use bson.ObjectId 
func genProductId() string {
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

func getProducts(ids []string) (products []*api.Product, err error) {
	idBsonMs := []bson.M{}
	for _, id := range ids {
		idBsonMs = append(idBsonMs, bson.M{"Id": id})
	}
	productMaps := []bson.M{}
	err = db.C(productTN).Find(bson.M{"$or": idBsonMs}).All(&productMaps)
	if err != nil {
		return
	}

	products = []*api.Product{}
	for _, productM := range productMaps {
		products = append(products, newProduct(productM))
	}

	return
}

func newProduct(productMap map[string]interface{}) (product *api.Product) {
	product = &api.Product{}

	// TODO codes below are terrible, refactor needed.
	if productMap["Name"] != nil {
		product.Name = productMap["Name"].(string)
	}
	if productMap["PhotoLink"] != nil {
		product.PhotoLink = productMap["PhotoLink"].(string)
	}
	if productMap["Price"] != nil {
		product.Price = productMap["Price"].(int64)
	}
	if productMap["Id"] != nil {
		product.Id = productMap["Id"].(string)
	}

	return
}
