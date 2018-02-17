package services

import (
	"gopkg.in/mgo.v2"
	"goipmserver/settings"
)

var (
	mgoSession     *mgo.Session
	mongoDatabase string
	mongoHost string
)

func init() {
	mongoDatabase = settings.Get().MongoDatabase
	mongoHost = settings.Get().MongoHost
}

func getSession () *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongoHost)
		if err != nil {
			panic(err) // no, not really
		}
	}
	return mgoSession.Clone()
}

func withCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(mongoDatabase).C(collection)
	return s(c)
}

func SearchCollection (collection string, q interface{}, skip int, limit int) (searchResults []interface{}, searchErr string) {
	if !Exist(collection) {
		return nil, "Invalid collection name : "+collection
	}

	query := func(c *mgo.Collection) error {
		fn := c.Find(q).Skip(skip).Limit(limit).All(&searchResults)

		if limit < 0 {
			fn = c.Find(q).Skip(skip).All(&searchResults)
		}
		return fn
	}
	search := func() error {
		return withCollection(collection, query)
	}
	err := search()
	if err != nil {
		searchErr = "Database Error: " + err.Error()
	}
	return searchResults, searchErr
}

func InsertCollection (collection string, data interface{}) (insertResult interface{}, insertError string) {
	if !Exist(collection) {
		return nil, "Invalid collection name : "+collection
	}

	err, ok := Validate(collection, data)
	if !ok {
		return nil, err
	}

	cmd := func(c *mgo.Collection) error {
		fn := c.Insert(data)
		return fn
	}

	insert := func() error {
		return withCollection(collection, cmd)
	}

	ierr := insert()
	if ierr != nil {
		insertError = "Database Error: " + ierr.Error()
	}
	return insertResult, insertError
}