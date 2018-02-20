package models

import (
	"errors"
)

type Resource struct {
	Id					string `bson:"_id" json:"_id"`
	System 				Base
	Code				string
	FirstName			string
	LastName			string
	Initial				string
	Address				Address
	WorkHoursByDay		float64
	WorkHoursByWeek		float64
	Cost				float64
	Company				string
}

const ResourceCollectionName = "Resources"

func (resource *Resource) Validate() error {
	if resource.Id == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func ValidateResource(data interface{}) (out interface{}, err error) {
	var audit Resource
	err = SetStruct(data, &audit)
	if err != nil {
		return audit, err
	}
	err = audit.Validate()
	return audit, err
}