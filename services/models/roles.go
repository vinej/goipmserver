package models

import (
	"errors"
)

type Role struct {
	Id					string `bson:"_id" json:"_id"`
	System 				Base
	Name				string
	Desc				string
	ExpectedCostByHour	float64
	RateByHour			float64
}

const RoleCollectionName = "roles"

func (role *Role) Validate() error {
	if role.Id == "" {
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

func PatchRole(data interface{}, patches []Patches) (out interface{}, err error) {
	var role Role
	err = SetStruct(data, &role)
	if err != nil {
		return role, err
	}
	err = PatchStruct(role, patches)
	return role, err
}