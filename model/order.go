package model

import (
	"api"
	"strings"
	"labix.org/v2/mgo/bson"
)

// TN => table name
var orderTN = "orders"

func (model Model) PutOrder(date string, email string, productIds []string) (order *api.Order, err error) {
	// TODO => product.rb:14
	newOrder := make(map[string]interface{})
	orderId := genOrderId(date, email)
	newOrder["Date"] = date
	newOrder["Email"] = email
	newOrder["Products"] = productIds
	newOrder["Id"] = orderId
	
	err = model.put(orderTN, newOrder, "Id", orderId)
	if err != nil {
		return
	}
	
	// TODO => product.rb:27
	order = &api.Order{}
	orderMap := map[string]interface{}{}
	err = db.C(userTN).Find(bson.M{"Email": email}).One(orderMap)
	if err != nil {
		return
	}
	order.User = getUser(orderMap["Email"].(string))
	order.Date = orderMap["Date"].(string)
	order.Products, err = getProducts(orderMap["Products"].([]string))
	if err != nil {
		return
	}
	
	return
}

func (model Model) RemoveOrder(date string, email string) {
	model.remove(orderTN, "Id", genOrderId(date, email))
	return
}

func genOrderId(date string, email string) string {
	return strings.Join([]string{date, email}, ":")	
}