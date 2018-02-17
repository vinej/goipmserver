package services

import (
	"strings"
	"goipmserver/services/models"
)

type validator func (interface{}) (string, bool)

var (
	allValidators map[string]validator
	allCollections = "companies:  audits: activities: orders: plans: resources: roles: users:"
)

func init() {
	allValidators = make(map[string]validator)
	allValidators[models.CompanyCollectionName] = models.ValidateCompany
	allValidators[models.UserCollectionName] = models.ValidateUser
}

func Validate(collection string, data interface{}) (string, bool) {
	var v = allValidators[collection]
	return v(data)
}

func Exist(collection string) bool {
	return strings.Contains(allCollections, collection+":")
}

