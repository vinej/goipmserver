package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"gopkg.in/mgo.v2/bson"
)

// Timestamp is a time but it serializes to ISO8601 format with millis
type Timestamp struct {
	time.Time
}

// ISO8601 format to millis instead of to nanos
const RFC3339Millis = "2006-01-02T15:04:05.000Z07:00"

func (t Timestamp) String() string {
	return t.Format(RFC3339Millis)
}

// ParseTimestamp parses a string that represents an ISO8601 time or a unix epoch
func ParseTimestamp(data string) (Timestamp, error) {
	d := time.Now().UTC()
	if data != "now" {
		// fmt.Println("we should try to parse")
		dd, err := time.Parse(RFC3339Millis, data)
		if err != nil {
			dd, err = time.Parse(time.RFC3339, data)
			if err != nil {
				dd, err = time.Parse(time.RFC3339Nano, data)
				if err != nil {
					if data == "" {
						data = "0"
					}
					t, err := strconv.ParseInt(data, 10, 64)
					if err != nil {
						return Timestamp{}, err
					}
					dd = time.Unix(0, t*int64(time.Millisecond))
				}
			}
		}
		d = dd
	}
	return Timestamp{Time: d.UTC()}, nil
}

// GetBSON customizes the bson serialization for this type
func (t Timestamp) GetBSON() (interface{}, error) {
	return t.Time, nil
}

// SetBSON customizes the bson serialization for this type
func (t *Timestamp) SetBSON(raw bson.Raw) error {
	var ts interface{}
	if err := raw.Unmarshal(&ts); err != nil {
		return err
	}
	switch ts.(type) {
	case time.Time:
		*t = Timestamp{Time: ts.(time.Time).UTC()}
		return nil
	case string:
		tss := ts.(string)
		tt, err := ParseTimestamp(tss)
		if err != nil {
			return err
		}
		*t = tt
		return nil
	case int64:
		*t = Timestamp{time.Unix(0, ts.(int64)*int64(time.Millisecond)).UTC()}
		return nil
	case float64:
		*t = Timestamp{time.Unix(0, int64(ts.(float64))*int64(time.Millisecond)).UTC()}
		return nil
	}

	return fmt.Errorf("couldn't convert bson data (%T) %s to a Timestamp", ts, ts)
}

// MarshalText implements the text marshaller interface
func (t Timestamp) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText reads this timestamp from a string value
func (t *Timestamp) UnmarshalText(data []byte) error {
	var value interface{}
	json.Unmarshal(data, &value)

	switch value.(type) {
	case string:
		v := value.(string)
		if v == "" {
			return nil
		}
		d, err := ParseTimestamp(v)
		if err != nil {
			return err
		}
		*t = d
	case float64:
		*t = Timestamp{time.Unix(0, int64(value.(float64))*int64(time.Millisecond)).UTC()}
	default:
		return fmt.Errorf("couldn't convert json from (%T) %s to a time.Time", value, data)
	}
	return nil
}

// UnmarshalJSON implements the json unmarshaller interface
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var value interface{}
	json.Unmarshal(data, &value)

	switch value.(type) {
	case string:
		v := value.(string)
		if v == "" {
			return nil
		}
		d, err := ParseTimestamp(v)
		if err != nil {
			return err
		}
		*t = d
	case float64:
		*t = Timestamp{time.Unix(0, int64(value.(float64))*int64(time.Millisecond)).UTC()}
	default:
		return fmt.Errorf("Couldn't convert json from (%T) %s to a time.Time", value, data)
	}
	return nil
}