package model

import (
	"github.com/sunfmin/batchbuy/api"
	"labix.org/v2/mgo/bson"
	// "fmt"
)

var userTN = "users"
var userCol = db.C(userTN)

type User struct {
	Id bson.ObjectId "_id"
	Name       string
	Email      string
	AvatarLink string
}

// TODO this design seems unnecesary, and make it more troublesome to use User#Put
type UserInput api.UserInput

func (user *User) Put(email string, input UserInput) (err error) {
	count, err := userCol.Find(M{"email": email}).Count()
	if count == 0 {
		userCol.Insert(input)
	} else {
		userCol.Update(M{"email": input.Email}, &input)
	}
	
	userCol.Find(M{"email": input.Email}).One(user)
	
	return
}

func (User) Remove(email string) (err error) {
	err = userCol.Remove(M{"email": email})
	
	return
}

func (user User) ToApi() (apiUser *api.User) {
	apiUser = &api.User{}
	apiUser.Name = user.Name
	apiUser.Email = user.Email
	apiUser.AvatarLink = user.AvatarLink
	
	return
}