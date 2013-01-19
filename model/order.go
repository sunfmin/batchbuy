package model

import (
	"api"
	"strings"
	// "labix.org/v2/mgo"
)

// TN => table name
var orderTN = "orders"

func (model Model) PutOrder(date string, email string, productIds []string) (order *api.Order, err error) {
	// // TODO => product.rb:14
	// newOrder := make(map[string]interface{})
	// newOrder["Email"] = input.Date
	// newOrder["Name"] = input.Email
	// // newUser["AvatarLink"] = input.AvatarLink
	// 
	// err = model.put(userTN, newUser, "Email", email)
	// if err != nil {
	// 	return
	// }
	// 
	// // TODO => product.rb:27
	// user = &api.User{}
	// userMap := map[string]interface{}{}
	// err = db.C(userTN).Find(bson.M{"Email": email}).One(userMap)
	// if err != nil {
	// 	return
	// }
	// user.Email = userMap["Email"].(string)
	// user.Name = userMap["Name"].(string)
	// user.AvatarLink = userMap["AvatarLink"].(string)
	
	// order_collection := db.C("order")
	// order = api.Order{date: date, email: email, productsIds: productIds}
	// order_collection.Insert()
	
	return
}

func (model Model) RemoveOrder(date string, email string) {
	// orderC := db.C(orderTN)
	// err = orderC.Remove(bson.M{"Date": date, "Email": email})
	return
}

func genOrderId(date string, email string) string {
	return strings.Join([]string{date, email}, ":")	
}