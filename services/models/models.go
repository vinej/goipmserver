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