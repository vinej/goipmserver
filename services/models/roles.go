package models

import "errors"

type Role struct {
	Name				string
	IsSync				bool
	ExpectedCostByHour	string
	UpdatedDate			string
	CreatedDate			string
	UpdatedBy			string
	UpdatedByOnServer	string
	IsNew				bool
	Id					string
	IsDeleted			bool
	Desc				string
	Version				int
	RateByHour			float64
	CreatedBy			string
	Order				float64
}

const RoleectionName = "plans"

func (role *Role) Validate() error {
	if role.Id == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func VaidateRolen(data interface{}) (out interface{}, err error) {
	var role Role
	err = SetStruct(data, &role)
	if err != nil {
		return role, err
	}
	err = role.Validate()
	return role, err
}