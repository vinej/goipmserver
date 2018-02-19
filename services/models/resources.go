package models

import "errors"

type Resource struct {
	updatedDateOnServer	string
	IsSync				bool
	Email				string
	UpdatedDate			string
	CreatedDate 		string
	Address				Address
	LastName			string
	Initial				string
	WorkHoursByDay		float64
	Telephone			string
	UpdatedBy			string
	UpdatedByOnServer	string
	IsNew				bool
	Id					string
	IsDeleted			bool
	Cost				float64
	Code				string
	WorkHoursByWeek		float64
	Version				int
	FirstName			string
	CreatedBy			string
	Order				float64
	Company				string
}

const ResourceCollectionName = "Resources"

func (resource *Resource) Validate() error {
	if resource.Id == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func VaidateResource(data interface{}) (out interface{}, err error) {
	var audit Resource
	err = SetStruct(data, &audit)
	if err != nil {
		return audit, err
	}
	err = audit.Validate()
	return audit, err
}