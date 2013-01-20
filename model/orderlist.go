package model

import (
	"api"
	"labix.org/v2/mgo/bson"
)

func (model Model) OrderListOfDate(date string) (orders []*api.Order, err error) {
	orderC := db.C("order")
	orders = []*api.Order{}
	err = orderC.Find(bson.M{"Date": date}).All(orders)
	
	return
}