// most tests here are hardcoded or ad-hoc, need to find a good way to fulfill these tests well.
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
	{orderTN, "Id", bson.M{"Email": "test@test.com", "Date": "2013-1-1"}, "PutUser"},
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
		case orderTN:
			testModel.PutOrder("2013-1-1", "test@test.com", []string{"no_1", "no_2"})
		}
		
		count, _ := dbC.Find(ft.testIdBsonM).Count()
		if count != 1 {
			t.Error("Model#%s Can't successfully create %s", ft.funcName, ft.col)
		}

		// clean up db
		// dbC.RemoveAll(ft.testIdBsonM)
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
		bson.M{
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
		bson.M{
			"Email": "test email",
			"Name": "test user",
		},
		"PutUser",
	},
	{
		orderTN,
		"2013-1-1:test@test.com",
		bson.M{"Id": "2013-1-1:test@test.com"},
		bson.M{
			"Email": "test@test.com",
			"Date": "2013-1-1",
			"Id": "2013-1-1:test@test.com",
		},
		"PutOrder",
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
			result := bson.M{}
			dbC.Find(ft.testIdBsonM).One(&result)
			
			failure = result["Name"].(string) != "test_produdct"	// TODO use interface here to help refactor tests
		case userTN:
			userInput := api.UserInput{Name: "new name", Email: "test email"}
			testModel.PutUser("test email", userInput)
			
			result := bson.M{}
			dbC.Find(ft.testIdBsonM).One(&result)
			
			failure = result["Name"].(string) != "new name"
		case orderTN:
			newProds := []string{"no_1", "no_2"}
			testModel.PutOrder("2013-1-1", "test@test.com", newProds)
			
			result := bson.M{}
			dbC.Find(ft.testIdBsonM).One(&result)
			
			if len(result["Products"].([]interface{})) != len(newProds) || 
				result["Products"].([]interface{})[0].(string) != "no_1" || 
				result["Products"].([]interface{})[1].(string) != "no_2" {
				failure = true
			}
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
	{
		orderTN,
		"2013-1-1:bom.d.van@gmail.com",
		bson.M{"Id": "2013-1-1:bom.d.van@gmail.com"},
		map[string]interface{}{
			"Id": "2013-1-1:bom.d.van@gmail.com",
			"Email": "bom.d.van@gmail.com",
			"Date": "2013-1-1",
			"Products": "",
		},
		"RemoveOrder",
	},
}

func TestRemove(t *testing.T) {
	for _, ft := range rmflagtests {
		dbC := db.C(ft.col)
		dbC.RemoveAll(ft.testIdBsonM)
		
		dbC.Insert(ft.testData)
		
		testModel := Model{}
		switch ft.col {
		case productTN:
			testModel.RemoveProduct(ft.testId)
		case userTN:
			testModel.RemoveUser(ft.testId)
		case orderTN:
			testModel.RemoveOrder(ft.testData["Date"].(string), ft.testData["Email"].(string))
		}
		
		count, _ := dbC.Find(ft.testIdBsonM).Count()
		if count != 0 {
			t.Errorf("Model#%s Can't remove %s properly", ft.funcName, ft.col)
		}
	}
}

func TestGetProducts(t *testing.T) {
	productC := db.C(productTN)
	productC.RemoveAll(bson.M{})
	productC.Insert(bson.M{"Id": "no_1"}, bson.M{"Id": "no_2"}, bson.M{"Id": "no_3"})
	
	products, _ := getProducts([]string{"no_1", "no_2", "no_3"})
	
	if len(products) != 3 {
		t.Errorf("Expeect get 3 products, only %d received", len(products))
	}
	
	productC.RemoveAll(bson.M{})
}