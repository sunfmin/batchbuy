package model

import (
	"api"
	"labix.org/v2/mgo/bson"
)

var userTN = "users"

func (model Model) PutUser(email string, input api.UserInput) (user *api.User, err error) {
	// TODO => product.rb:14
	newUser := make(map[string]interface{})
	newUser["Email"] = input.Email
	newUser["Name"] = input.Name
	// newUser["AvatarLink"] = input.AvatarLink
	
	err = model.put(userTN, newUser, "Email", email)
	if err != nil {
		return
	}
	
	// TODO => product.rb:27
	user = &api.User{}
	userMap := map[string]interface{}{}
	err = db.C(userTN).Find(bson.M{"Email": email}).One(userMap)
	if err != nil {
		return
	}
	user.Email = userMap["Email"].(string)
	user.Name = userMap["Name"].(string)
	// user.AvatarLink = userMap["AvatarLink"].(string)
	
	return
}

func (model Model) RemoveUser(email string) (err error) {
	err = model.remove(userTN, "Email", email)
	return
}

func getUser(email string) (user *api.User) {
	db.C(userTN).Find(bson.M{"Email": email}).One(user)
	return
}