package models

import (
	"errors"
)

const UserCollectionName = "users"

type User struct {
	Id			string `bson:"_id" json:"_id"`
	System 		Base
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

func PatchUser(data interface{}, patches []Patches) (out interface{}, err error) {
	var user User
	err = SetStruct(data, &user)
	if err != nil {
		return user, err
	}
	err = PatchStruct(user, patches)
	return user, err
}

