package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type ActivityHistory struct {
	UUID				bson.ObjectId `bson:"_id,omitempty"`
	System 				Base
	StartDate			Timestamp
	EndDate				Timestamp
	ResourceStart		string
	ResourceEnd			string
	WorkFlow			string
	Duration			float64
}

const ActivityHistoryCollectionName = "activities_histories"

func (activityHistory *ActivityHistory) Validate() error {
	if activityHistory.System.Id == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func VaidateActivityHistory(data interface{}) (out interface{}, err error) {
	var activityHistory ActivityHistory
	err = SetStruct(data, &activityHistory)
	if err != nil {
		return activityHistory, err
	}
	err = activityHistory.Validate()
	return activityHistory, err
}

