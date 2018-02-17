package models

import (
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type Company struct {
	ID			bson.ObjectId `bson:"_id,omitempty"`
	Name 		string
}

const CompanyCollectionName = "companies"

func ValidateCompany(data interface{}) (string, bool) {
	var cie Company
	byteData, err := json.Marshal(data)
	if err != nil {
		return err.Error(), false
	}
	err = json.Unmarshal(byteData, &cie)
	if err != nil {
		return err.Error(), false
	}
	if cie.Name == "" {
		return "invalid field content <name>", false
	}
	return "ok", true
}
