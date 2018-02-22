package models

import (
	"errors"
)

type Plan struct {
	Id					string `bson:"_id" json:"_id"`
	System 				Base
	Code				string
	Name				string
	Desc				string
	ExpectedMargin		float64
	ScheduleStartDate	Timestamp
	ScheduleEndDate		Timestamp
	InitialBudget		float64
	ContingencyBudget	float64
	Risk				float64
	Status				string
	Timezone			float64
	Project				string
}

const PlanCollectionName = "plans"

func (plan *Plan) Validate() error {
	if plan.Id == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func ValidatePlan(data interface{}) (out interface{}, err error) {
	var plan Plan
	err = SetStruct(data, &plan)
	if err != nil {
		return plan, err
	}
	err = plan.Validate()
	return plan, err
}

func PatchPlan(data interface{}, patches []Patches) (out interface{}, err error) {
	var plan Plan
	err = SetStruct(data, &plan)
	if err != nil {
		return plan, err
	}
	err = PatchStruct(plan, patches)
	return plan, err
}
