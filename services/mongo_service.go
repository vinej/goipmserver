package services

import (
	"gopkg.in/mgo.v2"
	"goipmserver/settings"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"goipmserver/services/models"
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

func SearchCollection (collection string, q interface{}, skip int, limit int) (results []interface{}, err error) {
	if !Exist(collection) {
		return nil, errors.New("Invalid collection name : "+collection)
	}

	query := func(c *mgo.Collection) error {
		fn := c.Find(q).Skip(skip).Limit(limit).All(&results)

		if limit < 0 {
			fn = c.Find(q).Skip(skip).All(&results)
		}
		return fn
	}
	search := func() error {
		return withCollection(collection, query)
	}
	serr := search()
	if serr != nil {
		err = errors.New("Database Error: " + serr.Error())
	}
	return results, err
}

func InsertCollection (collection string, data interface{}) (results []interface{}, err error) {
	if !Exist(collection) {
		return nil, errors.New("Invalid collection name : "+collection)
	}

	out, err := Validate(collection, data)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var baseSystem models.BaseId
	serr := models.SetStruct(data, &baseSystem)
	if serr != nil {
		return  nil, errors.New("Database Error: " + serr.Error())
	}

	cmd := func(c *mgo.Collection) error {
		fn := c.Insert(out)
		return fn
	}

	insert := func() error {
		return withCollection(collection, cmd)
	}

	ierr := insert()
	if ierr != nil {
		return  nil, errors.New("Database Error: " + ierr.Error())
	}

	return SearchCollection(collection, bson.M{ "_id": baseSystem.Id },0,1)
}

func UpdateCollection (collection string, data interface{}) (results []interface{}, err error) {
	if !Exist(collection) {
		return nil,errors.New("Invalid collection name : "+collection)
	}

	out, err := Validate(collection, data)
	if err != nil {
		return nil,errors.New(err.Error())
	}

	// get the old document
	var baseSystem models.BaseId
	serr := models.SetStruct(data, &baseSystem)
	if serr != nil {
		return  nil, errors.New("Database Error: " + serr.Error())
	}

	// MERGE/CONFLIC MANAGEMENT

	cmd := func(c *mgo.Collection) error {
		fn := c.Update(bson.M{ "_id": baseSystem.Id}, out)
		return fn
	}

	update := func() error {
		return withCollection(collection, cmd)
	}

	ierr := update()
	if ierr != nil {
		return  nil,errors.New("Database Error: " + ierr.Error())
	}

	return SearchCollection(collection, bson.M{ "_id": baseSystem.Id},0,1)
}

func DeleteCollection (collection string, data interface{}) error {
	if !Exist(collection) {
		return errors.New("Invalid collection name : "+collection)
	}

	// get the old document
	var baseSystem models.BaseId
	serr := models.SetStruct(data, &baseSystem)
	if serr != nil {
		return  errors.New("Database Error: " + serr.Error())
	}

	cmd := func(c *mgo.Collection) error {
		fn := c.Remove(bson.M{ "_id": baseSystem.Id})
		return fn
	}

	remove := func() error {
		return withCollection(collection, cmd)
	}

	ierr := remove()
	if ierr != nil {
		// note:if the doc is already deleted must return OK
		if ierr.Error() == "not found" {
		    return nil
	    }
		return  errors.New("Database Error: " + ierr.Error())
	}
	return nil
}