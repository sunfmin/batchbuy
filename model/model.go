package model

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Model struct {}

func startSession() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")	
	if err != nil {
		panic(err)
	}
	// defer session.Close() // TODO figure out will not closing session causse any problem
	
	return session
}

var db = startSession().DB("low_tea_at_the_plant")

/*
  	`docValues` must contains `idVal` beforehand
  	`wouldBeIdVal` is used for create new doc, this design is due to the generation of id of product
  	documents is different the rest of other collections
*/
func (model Model) put(collection string, docValues map[string]interface{}, idKey string, wouldBeIdVal string) (err error) {
	dbC := db.C(collection)
	
	docId := bson.M{idKey: docValues[idKey]}
	query := dbC.Find(docId)
	count, err := query.Count()
	if err != nil {
		return
	}
	
	if (count == 0) {
		docValues[idKey] = wouldBeIdVal
		
		// TODO figure out why `Insert` needs a pointer
		// TODO figure out `Inster` turn `struct field` into downcase but not to `map`
		err = dbC.Insert(&docValues)
		if err != nil {
			return
		}
	} else {
		dbC.Update(docId, docValues)
	}
	
	return
}

func (model Model) remove(collection string, idKey string, idVal string) (err error) {
	dbC := db.C(collection)
	err = dbC.Remove(bson.M{idKey: idVal})
	return
}