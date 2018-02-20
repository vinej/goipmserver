package models

import (
	"errors"
)

type Audit struct {
	Id					string `bson:"_id" json:"_id"`
	System 				Base
	ObjectName			string
	ObjectId			string
	ObjectDate			string
	AuditAction			string
}

const AuditCollectionName = "audits"

func (audit *Audit) Validate() error {
	if audit.Id == "" {
		return errors.New("invalid field content <name>")
	}
	return nil
}

func VaidateAudit(data interface{}) (out interface{}, err error) {
	var audit Audit
	err = SetStruct(data, &audit)
	if err != nil {
		return audit, err
	}
	err = audit.Validate()
	return audit, err
}