package models

import (
	"errors"
)

type Document interface {
	Validate() error
}

type Activity struct {
	Id					string `bson:"_id" json:"_id"`
	System 				Base
	Code				string
	Name 				string
	FixeStartDate 		Timestamp
	TotalDuration 		float64
	ExpectedDuration	float64
	Resource 			string
	BackupResource		string
	Role 				string
	Plan				string
	WorkFlow 			string
}

const ActivityCollectionName = "activities"

func (activity *Activity) Validate() error {
	if activity.Id == "" {
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

func PatchActivity(data interface{}, patches []Patches) (out interface{}, err error) {
	var activity Activity
	err = SetStruct(data, &activity)
	if err != nil {
		return activity, err
	}
	err = PatchStruct(activity, patches)
	return activity, err
}