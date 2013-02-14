package model

import (
	"fmt"
	"testing"
)

// StartConnectDb("test")

func TestCreateUser(t *testing.T) {
	userCol.RemoveAll(M{})

	userInput := UserInput{"test", "test@test.com", "link to avatar"}
	user := User{}
	user.Put("non-exited user", userInput)

	result := User{}
	userCol.Find(M{"email": user.Email}).One(&result)

	if user != result {
		t.Error("Can't Create New User")
	}

	userCol.RemoveAll(M{})
}

func TestUpdateUser(t *testing.T) {
	userCol.RemoveAll(M{})
	userInput := UserInput{"test", "test@test.com", "link to avatar"}
	user := User{}
	user.Put("non-exited user", userInput)

	userInput.Name = "new name"
	user.Put(user.Email, userInput)

	result := User{}
	userCol.FindId(user.Id).One(&result)

	if user != result && user.Name != "new name" {
		fmt.Println("Expect: ", user)
		fmt.Println("Got: ", result)
		t.Error("Can't Update User")
	}

	userCol.RemoveAll(M{})
}

// func TestAvaliableProducts(t *testing.T) {
//     initDbForProductTests()
//     userInput := UserInput{"test", "test@test.com", "link to avatar"}
//     user := User{}
//     user.Put("non-exited user", userInput)
// }
