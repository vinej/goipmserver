package models

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

const UserCollectionName = "users"

type User struct {
	ID		bson.ObjectId `bson:"_id,omitempty"`
	User	 	string
	Email       string
	Password 	string
	Admin 		bool
	RegisterOn  string `json:"register_on"`
}

func GetUser(data interface{}) (user User, err string) {
	var tuser User
	terr := SetStruct(data, &tuser)
	fmt.Println(tuser.User)
	return tuser, terr
}

func ValidateUser(data interface{}) (string, bool) {
	user, err := GetUser(data)
	if err != "" {
		return err, false
	}
	if user.User == "" {
		return "user can't be empty", false
	}
	return "ok", true
}

