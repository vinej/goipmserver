package models

import (
	"errors"
)

type Activity struct {
	IsSync  			bool
	TotalDuration 		float64
	CreatedDate			string
	BackupResource		string
	IsNew				bool
	Id					string
	Code				string
	MinType				float64
	MaxType				float64
	TypeInfo			string
	Plan				string
	UpdatedDateOnServer	string
	Name 				string
	StartDate 			string
	UpdatedDate 		string
	WorkFlow 			string
	FixeStartDate 		string
	IncrType 			string
	Role 				string
	UpdatedBy 			string
	UpdatedByOnServer 	string
	ExpectedDuration	float64
}

const ActivityCollectionName = "activities"

func (activity *Activity) Validate() error {
	if activity.Name == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func VaidateActivity(data interface{}) (out interface{}, err error) {
	var activity Activity
	err = SetStruct(data, &activity)
	if err != nil {
		return activity, err
	}
	err = activity.Validate()
	return activity, err
}
