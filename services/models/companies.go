package models

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Company struct {
	ID					bson.ObjectId `bson:"_id,omitempty"`
	Name 				string
	Code 				string
	UpdatedDateOnServer string
	IsSync				bool
	UpdatedDate 		string
	CreatedDate 		string
	Address 			Address
	Type 				string
	UpdatedBy 			string
	UpdatedByOnServer 	string
	IsNew 				bool
	IsDeleted 			bool
	Version 			int
	CreatedBy 			string
	Order 				float64
}

const CompanyCollectionName = "companies"

func (company *Company) Validate() error {
 	if company.Name == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func VaidateCompany(data interface{}) (out interface{}, err error) {
	var company Company
	err = SetStruct(data, &company)
	if err != nil {
		return company, err
	}
	err = company.Validate()
	return company, err
}
