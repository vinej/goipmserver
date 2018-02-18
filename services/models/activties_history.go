package models

import (
	"errors"
)

type ActivityHistory struct {
	UpdatedDateOnServer		string
	IsSync					bool
	StartDate				string
	UpdatedDate				string
	CreatedDate				string
	ResourceEnd				string
	WorkFlow				string
	UpdatedBy				string
	UpdatedByOnServer		string
	IsNew					bool
	Id						string
	IsDeleted				bool
	ResourceStart			string
	EndDate					string
	Version					int
	Duration				float64
	CreatedBy				string
	Order					float64
}

const ActivityHistoryCollectionName = "activities"

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

