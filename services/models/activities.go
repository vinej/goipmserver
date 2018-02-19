package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Activity struct {
	UUID				bson.ObjectId `bson:"_id,omitempty"`
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
	if activity.System.Id == "" {
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
