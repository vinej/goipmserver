package models

import (
	"encoding/json"
)

func SetStruct(data interface{}, v interface{}) error {
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err2 := json.Unmarshal(byteData, &v)
	if err2 != nil {
		return err2
	}

	return nil
}