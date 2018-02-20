package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Role struct {
	UUID				bson.ObjectId `bson:"_id,omitempty"`
	System 				Base
	Name				string
	Desc				string
	ExpectedCostByHour	float64
	RateByHour			float64
}

const RoleCollectionName = "roles"

func (role *Role) Validate() error {
	if role.System.Id == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func ValidateRole(data interface{}) (out interface{}, err error) {
	var role Role
	err = SetStruct(data, &role)
	if err != nil {
		return role, err
	}
	err = role.Validate()
	return role, err
}