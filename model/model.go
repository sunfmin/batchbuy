package model

import (
	"labix.org/v2/mgo"
	"encoding/hex"
	// "labix.org/v2/mgo/bson"
)

type Model struct {}
type M map[string]interface{}

// var session mgo.Session
func Init() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")	
	if err != nil {
		panic(err)
	}
	// defer session.Close() // TODO: figure out will not closing session causse any problem
	
	return session
}

// func End() {
// 	session.Close()
// }

var db = Init().DB("low_tea_at_the_plant")

// Ported from mgo/bson for the sake of not knowing how to update mgo
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

// func Model Error() string{
// 	return "Can't get that action done."
// }