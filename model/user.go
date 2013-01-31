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
    if err != nil {
        return
    }
	if count == 0 {
		err = userCol.Insert(input)
	} else {
		err = userCol.Update(M{"email": input.Email}, &input)
	}
    if err != nil {
        return
    }
	
	err = userCol.Find(M{"email": input.Email}).One(user)
    if err != nil {
        return
    }
	
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