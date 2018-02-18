package services

import (
	"strings"
	"goipmserver/services/models"
)

type validator func (interface{}) (interface{}, error)

var (
	allValidators map[string]validator
	allCollections = "companies:  audits: activities: orders: plans: resources: roles: users:"
)

func init() {
	allValidators = make(map[string]validator)
	allValidators[models.CompanyCollectionName] = models.VaidateCompany
	allValidators[models.UserCollectionName] = models.ValidateUser
}

func Validate(collection string, data interface{}) (out interface{}, err error) {
	var v = allValidators[collection]
	return v(data)
}

func Exist(collection string) bool {
	return strings.Contains(allCollections, collection+":")
}

