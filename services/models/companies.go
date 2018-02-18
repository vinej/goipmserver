package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Company struct {
	ID			bson.ObjectId `bson:"_id,omitempty"`
	Name 		string
}

const CompanyCollectionName = "companies"

func ValidateCompany(data interface{}) (string, bool) {
	var cie Company
	err := SetStruct(data, &cie)
	if err != "" {
		return err, false
	}
 	if cie.Name == "" {
		return "invalid field content <name>", false
	}
	return "ok", true
}
