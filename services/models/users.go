package models

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
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

func (user *User) Validate() error {
	if user.User == "" {
		return errors.New("user can't be empty")
	}
	return nil
}

func ValidateUser(data interface{}) (out interface{}, err error) {
	var user User
	err = SetStruct(data, &user)
	if err != nil {
		return user, err
	}
	err = user.Validate()
	return user, err
}

