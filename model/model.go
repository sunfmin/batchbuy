package model

import (
	"encoding/hex"

	"labix.org/v2/mgo"
)

type Model struct{}
type M map[string]interface{}

func Init() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
	// session, err := mgo.DialWithInfo(&mgo.DialInfo{
	// 	Addrs:    []string{"localhost:27017"},
	// 	Username: "qortexdev",
	// 	Password: "JT_PX5AQ@m55mNyud@JD",
	// })
	// println("Connecting to DB")
	// session, err := mgo.DialWithInfo(&mgo.DialInfo{
	// 	Addrs:    []string{"ds053419.mongolab.com:53419/heroku_app27877108"},
	// 	Username: "lowtea",
	// 	Password: "e34R@ht8U2DJei",
	// })
	if err != nil {
		panic(err)
	}

	println("Connecting DB Done")
	return session
}

func End() {
	session.Close()
}

var session = Init()
var db = session.DB("low_tea_at_the_plant")

// Borrow from mgo/bson for the sake of can't update mgo
// TODO: figure out how to update go remote package and remove this function
// IsObjectIdHex returns whether s is a valid hex representation of
// an ObjectId. See the ObjectIdHex function.
func isObjectIdHex(s string) bool {
	if len(s) != 24 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}
