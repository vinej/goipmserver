package models

import (
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	//"goipmserver/settings"
)

type Company struct {
	ID			bson.ObjectId `bson:"_id,omitempty"`
	Name 		string
}

const CompanyCollectionName = "companies"

func GetCompany(data interface{}) (c Company, err string) {
	var cie Company
	byteData, error := json.Marshal(data)
	if error != nil {
		return cie, error.Error()
	}
	error = json.Unmarshal(byteData, &cie)
	if error != nil {
		return cie, error.Error()
	}
	return cie, ""
}

func ValidateCompany(data interface{}) (string, bool) {
	cie, err := GetCompany(data)
	if err != "" {
		return err, false
	}
 	if cie.Name == "" {
		return "invalid field content <name>", false
	}
	return "ok", true
}
