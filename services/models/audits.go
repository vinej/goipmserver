package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Audit struct {
	UUID				bson.ObjectId `bson:"_id,omitempty"`
	System 				Base
	ObjectName			string
	ObjectId			string
	ObjectDate			string
	AuditAction			string
}

const AuditCollectionName = "audits"

func (audit *Audit) Validate() error {
	if audit.System.Id == "" {
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