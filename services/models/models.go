package models

import (
	"encoding/json"
	"gopkg.in/oleiade/reflections.v1"
	"strconv"
	"fmt"
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

type Conflict struct {
	FieldName     string
	ServerValue   string
	ClientValue   string
}

type BaseId struct {
	Id string `bson:"_id" json:"_id"`
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

func SetPatches(data interface{}, v []Patches) error {
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

func GetConflit(data interface{}, patches []Patches) ( results []Conflict, err error) {
	var conflits = []Conflict{}
	for _,p := range patches{
		value, _ := reflections.GetField(&data, p.Field)
		fmt.Sprintf("%v", value)
		c := Conflict{ FieldName : p.Ftype, ServerValue: fmt.Sprintf("%v", value), ClientValue: p.Value }
		conflits = append(conflits, c )
	}
	return conflits, nil
}

func PatchStruct(data interface{}, patches []Patches) error {
	var perr error
	for _,p := range patches{
		switch(p.Ftype) {
		case "string":
			perr = reflections.SetField(&data, p.Field, p.Value)
		case "boolean":
			bval, err := strconv.ParseBool(p.Value)
			if (err != nil) {
				return err
			}
			perr = reflections.SetField(&data, p.Field, bval)
		case "number":
			fval, err := strconv.ParseFloat(p.Value, 64)
			if (err != nil) {
				return err
			}
			perr = reflections.SetField(&data, p.Field, fval)
		case "integer":
			ival, err := strconv.Atoi(p.Value)
			if (err != nil) {
				return err
			}
			perr = reflections.SetField(&data, p.Field, ival)
		case "date":
			perr = reflections.SetField(&data, p.Field, p.Value)
		}
		if (perr != nil) {
			return perr
		}
	}
	return nil
}
