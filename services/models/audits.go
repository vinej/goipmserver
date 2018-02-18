package models

import "errors"

type Audit struct {
	UpdatedDateOnServer	string
	IsSync				bool
	UpdatedDate			string
	CreatedDate			string
	ObjectName			string
	UpdatedBy			string
	UpdatedByOnServer	string
	IsNew				bool
	Id					string
	IsDeleted			bool
	AuditAction			string
	Version				int
	ObjectId			string
	CreatedBy			string
	Order				int
	ObjectDate			string
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