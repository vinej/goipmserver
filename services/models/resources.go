package models

import "errors"

type Resource struct {
	Id					string
	Code				string
	FirstName			string
	LastName			string
	Email				string
	Company				string
	UpdatedDateOnServer	string
	IsSync				bool
	UpdatedDate			string
	CreatedDate 		string
	Address				Address
	Initial				string
	WorkHoursByDay		float64
	Telephone			string
	UpdatedBy			string
	UpdatedByOnServer	string
	IsNew				bool
	IsDeleted			bool
	Cost				float64
	WorkHoursByWeek		float64
	Version				int
	CreatedBy			string
	Order				float64
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