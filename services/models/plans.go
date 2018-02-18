package models

import "errors"

type Plan struct {
	updatedDateOnServer string
	Name				string
	IsSync				bool
	ExpectedMargin		float64
	UpdatedDate			string
	CreatedDate			string
	ScheduleEndDate		string
	Risk				float64
	ScheduleStartDate	string
	Project				string
	ContingencyBudget	float64
	UpdatedBy			string
	UpdatedByOnServer	string
	IsNew				bool
	Id					string
	IsDeleted			bool
	Desc				string
	Status				string
	Code				string
	Version				int
	InitialBudget		float64
	Timezone			float64
	CreatedBy			string
}

const PlanCollectionName = "plans"

func (plan *Plan) Validate() error {
	if plan.Id == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func VaidatePlan(data interface{}) (out interface{}, err error) {
	var plan Plan
	err = SetStruct(data, &plan)
	if err != nil {
		return plan, err
	}
	err = plan.Validate()
	return plan, err
}