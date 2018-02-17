package models

import (
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

/*
type User struct {
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
*/
const UserCollectionName = "users"

type User struct {
	ID		bson.ObjectId `bson:"_id,omitempty"`
	User	 	string
	Email       string
	Password 	string
	Admin 		bool
	RegisterOn  string `json:"register_on"`
}

func GetUser( data interface{}) (user User, err string) {
	var usr User
	byteData, error := json.Marshal(data)
	if error != nil {
		return usr, error.Error()
	}
	error = json.Unmarshal(byteData, &usr)
	if error != nil {
		return usr, error.Error()
	}
	return usr, ""
}

func ValidateUser(data interface{}) (string, bool) {
	usr, err := GetUser(data)
	if err != "" {
		return err, false
	}
	if usr.User == "" {
		return "user can't be empty", false
	}
	return "ok", true
}

