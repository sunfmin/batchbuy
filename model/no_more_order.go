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

func NoMoreOrderToday() error {
	date := time.Now()
	return noMoreOrderCol.Insert(bson.M{"date": time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())})
}

func IsNoMoreOrderToday() (bool, error) {
	date := time.Now()
	count, err := noMoreOrderCol.Find(bson.M{"date": time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())}).Count()

	return count > 0, err
}

func MakeMoreOrderToday() error {
	date := time.Now()
	_, err := noMoreOrderCol.RemoveAll(bson.M{"date": time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())})
	return err
}
