package services

// see : http://sosedoff.com/2016/07/16/golang-struct-tags.html
// for a more advence validation

import (
	"strings"
	"goipmserver/services/models"
)

type validator func (interface{}) (interface{}, error)
type patchor func (interface{}, []models.Patches) (interface{}, error)
var (
	allValidators map[string]validator
	allPatchors map[string]patchor
	allCollections = "companies:  audits: activities: activity_history: orders: plans: resources: roles: users:"
)

func init() {
	allValidators = make(map[string]validator)
	allValidators[models.CompanyCollectionName] = models.VaidateCompany
	allValidators[models.UserCollectionName] = models.ValidateUser
	allValidators[models.AuditCollectionName] = models.VaidateAudit
	allValidators[models.ActivityCollectionName] = models.VaidateActivity
	allValidators[models.PlanCollectionName] = models.ValidatePlan
	allValidators[models.ResourceCollectionName] = models.ValidateResource
	allValidators[models.RoleCollectionName] = models.ValidateRole
	allValidators[models.ActivityHistoryCollectionName] = models.VaidateActivityHistory

	allPatchors = make(map[string]patchor)
	allPatchors[models.CompanyCollectionName] = models.PatchCompany
	allPatchors[models.UserCollectionName] = models.PatchUser
	allPatchors[models.AuditCollectionName] = models.PatchAudit
	allPatchors[models.ActivityCollectionName] = models.PatchActivity
	allPatchors[models.PlanCollectionName] = models.PatchPlan
	allPatchors[models.ResourceCollectionName] = models.PatchResource
	allPatchors[models.RoleCollectionName] = models.PatchRole
	allPatchors[models.ActivityHistoryCollectionName] = models.PatchActivityHistory

}

func Patch(collection string, data interface{}, patches []models.Patches) (out interface{}, err error) {
	var v = allPatchors[collection]
	return v(data, patches)
}

func Validate(collection string, data interface{}) (out interface{}, err error) {
	var v = allValidators[collection]
	return v(data)
}

func Exist(collection string) bool {
	return strings.Contains(allCollections, collection+":")
}

