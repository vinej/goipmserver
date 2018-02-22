package models

import (
	"errors"
)

type ActivityHistory struct {
	Id					string `bson:"_id" json:"_id"`
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
	if activityHistory.Id == "" {
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

func PatchActivityHistory(data interface{}, patches []Patches) (out interface{}, err error) {
	var activityHistory ActivityHistory
	err = SetStruct(data, &activityHistory)
	if err != nil {
		return activityHistory, err
	}
	err = PatchStruct(activityHistory, patches)
	return activityHistory, err
}

