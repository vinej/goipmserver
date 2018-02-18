package models

import (
	"encoding/json"
)

func SetStruct(data interface{}, v interface{}) (err string) {
	byteData, error := json.Marshal(data)
	if error != nil {
		return error.Error()
	}
	error = json.Unmarshal(byteData, &v)
	if error != nil {
		return error.Error()
	}

	return ""
}