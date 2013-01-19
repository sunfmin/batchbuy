// most tests here are hardcoded or ad-hoc, need to find a good way to fulfill these tests.
// TODO refactor needed. Table Driven Tests seem not very suitable for situations below
package model

import (
	"testing"
	"api"
	"labix.org/v2/mgo/bson"
)

var createflagtests = []struct {
	col string
	idKey string
	testIdBsonM bson.M
	funcName string
}{
	{productTN, "Id", bson.M{"Name": "test_produdct"}, "PutProduct"},
	{userTN, "Email", bson.M{"Name": "test_user"}, "PutUser"},
}

func TestCreate(t *testing.T) {
	for _, ft := range createflagtests {
		// prepare db
		dbC := db.C(ft.col)
		dbC.RemoveAll(bson.M{ft.idKey: ""})
		dbC.RemoveAll(ft.testIdBsonM)
		
		testModel := Model{}
		switch ft.col {
		case productTN:
			productInput := api.ProductInput{Name: "test_produdct"}
			testModel.PutProduct("", productInput)
		case userTN:
			userInput := api.UserInput{Name: "test_user", Email: "email"}
			testModel.PutUser("email", userInput)
		}
		count, _ := dbC.Find(ft.testIdBsonM).Count()
		if count != 1 {
			t.Error("Model#%s Can't successfully create %s", ft.funcName, ft.col)
		}

		// clean up db
		dbC.RemoveAll(ft.testIdBsonM)
	}
}

var updateflagtests = []struct {
	col string
	testId string
	testIdBsonM bson.M
	testData map[string]interface{}
	funcName string
}{
	{
		productTN, 
		"test id", 
		bson.M{"Id": "test id"},
		map[string]interface{}{
			"Id": "test id",
			"Name": "test name",
			"PhotoLink": "test photolink",
			"Price": 100,
		},
		"PutProduct",
	},
	{
		userTN,
		"test email",
		bson.M{"Email": "test email"},
		map[string]interface{}{
			"Email": "tetst email",
			"Name": "test user",
		},
		"PutUser",
	},
}

func TestUpdate(t *testing.T) {
	for _, ft := range updateflagtests {
		dbC := db.C(ft.col)
		dbC.RemoveAll(ft.testIdBsonM)
		
		dbC.Insert(ft.testData)
		
		testModel := Model{}
		failure := false
		switch ft.col {
		case productTN:
			productInput := api.ProductInput{Name: "test_produdct"}
			testModel.PutProduct("test id", productInput)
			
			result := map[string]interface{}{}
			dbC.Find(ft.testIdBsonM).One(&result)
			
			failure = result["Name"].(string) != "test_produdct"
		case userTN:
			userInput := api.UserInput{Name: "new name", Email: "test email"}
			testModel.PutUser("test email", userInput)
			
			result := map[string]interface{}{}
			dbC.Find(ft.testIdBsonM).One(&result)
			
			failure = result["Name"].(string) != "new name"
		}
		
		if failure {
			t.Errorf("Model#%s Can't successfully update %s", ft.funcName, ft.col)
		}
		
		dbC.RemoveAll(ft.testIdBsonM)
	}
}

var rmflagtests = []struct {
	col string
	testId string
	testIdBsonM bson.M
	testData map[string]interface{}
	funcName string
}{
	{
		productTN, 
		"test id", 
		bson.M{"Id": "test id"},
		map[string]interface{}{
			"Id": "test id",
			"Name": "test name",
			"PhotoLink": "test photolink",
			"Price": 100,
		},
		"RemoveProduct",
	},
	{
		userTN,
		"test email",
		bson.M{"Email": "test email"},
		map[string]interface{}{
			"Email": "test email",
			"Name": "test name",
		},
		"RemoveUser",
	},
}

func TestRemove(t *testing.T) {
	for _, ft := range rmflagtests {
		dbC := db.C(ft.col)
		dbC.RemoveAll(ft.testIdBsonM)
		
		dbC.Insert(ft.testData)
		
		testModel := Model{}
		if ft.col == productTN {
			testModel.RemoveProduct(ft.testId)
		} else if ft.col == userTN {
			testModel.RemoveUser(ft.testId)
		} else if ft.col == orderTN {
			// testModel.RemoveOrder(ft.testId)
		}
		
		count, _ := dbC.Find(ft.testIdBsonM).Count()
		if count != 0 {
			t.Errorf("Model#%s Can't remove %s properly", ft.funcName, ft.col)
		}
	}
}