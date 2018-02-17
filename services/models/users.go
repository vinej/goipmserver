package models

import "gopkg.in/mgo.v2/bson"

/*
type User struct {
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
*/

type User struct {
	ID		bson.ObjectId `bson:"_id,omitempty"`
	UUID        string
	Username 	string `json:"user"`
	Email       string
	Password 	string
	Admin 		bool
	RegisterOn  string `json:"register_on"`
}

func ValidateUser(b bson.M) (string, bool) {
	if b["name"] != nil {
	}
	return "ok", true
}

