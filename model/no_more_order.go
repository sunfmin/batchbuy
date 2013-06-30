package model

import (

	// "github.com/sunfmin/batchbuy/api"
	// "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
	// "fmt"
)

var noMoreOrderTN = "no_more_orders"
var noMoreOrderCol = db.C(noMoreOrderTN)

type NoMoreOrder struct {
	Id   bson.ObjectId "_id"
	Date time.Time
}

func NoMoreOrderToday(date time.Time) error {
	// date := time.Now()
	return noMoreOrderCol.Insert(bson.M{"date": date})
}

func IsNoMoreOrderToday(date time.Time) (bool, error) {
	count, err := noMoreOrderCol.Find(bson.M{"date": date}).Count()

	return count > 0, err
}

func MakeMoreOrderToday(date time.Time) error {
	// date := time.Now()
	_, err := noMoreOrderCol.RemoveAll(bson.M{"date": getDayRangeCond(date)})
	return err
}
