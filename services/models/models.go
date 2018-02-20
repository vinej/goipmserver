package models

import (
	"encoding/json"
)

type Address struct {
	State 		string
	City 		string
	Email 		string
	Street 		string
	PostalCode 	string
	Phone 		string
	PhoneHome 	string
	EmailHome 	string
	Country 	string
}

type Base struct {
	Id					string
	CreatedDate			Timestamp
	CreatedBy			string
	UpdatedDate 		Timestamp
	UpdatedBy 			string
	UpdatedDateOnServer	Timestamp
	UpdatedByOnServer 	string
	IsNew				bool
	isDeleted			bool
	IsSync  			bool
	Order				float64
	Version				int
}

type BaseSystem struct {
	System Base
}


func SetStruct(data interface{}, v interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err2 := json.Unmarshal(byteData, &v)
	if err2 != nil {
		return err2
	}

	return nil
}