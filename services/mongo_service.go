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

func InsertCollection (collection string, data interface{}) error {
	if !Exist(collection) {
		return errors.New("Invalid collection name : "+collection)
	}

	out, err := Validate(collection, data)
	if err != nil {
		return errors.New(err.Error())
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
		return  errors.New("Database Error: " + ierr.Error())
	}
	return nil
}

func UpdateCollection (collection string, data interface{}) error {
	if !Exist(collection) {
		return errors.New("Invalid collection name : "+collection)
	}

	out, err := Validate(collection, data)
	if err != nil {
		return errors.New(err.Error())
	}

	// get the old document
	var baseSystem models.BaseSystem
	serr := models.SetStruct(data, &baseSystem)
	if serr != nil {
		return  errors.New("Database Error: " + serr.Error())
	}

	searchResult, searchError:= SearchCollection(collection, bson.M{ "system.id": baseSystem.System.Id},0,1)
	if searchError != "" {
		return  errors.New("Database Error: " + searchError)
	}

	cmd := func(c *mgo.Collection) error {
		fn := c.Update(searchResult[0], out)
		return fn
	}

	update := func() error {
		return withCollection(collection, cmd)
	}

	ierr := update()
	if ierr != nil {
		return  errors.New("Database Error: " + ierr.Error())
	}
	return nil
}

func DeleteCollection (collection string, data interface{}) error {
	if !Exist(collection) {
		return errors.New("Invalid collection name : "+collection)
	}

	// get the old document
	var baseSystem models.BaseSystem
	serr := models.SetStruct(data, &baseSystem)
	if serr != nil {
		return  errors.New("Database Error: " + serr.Error())
	}

	cmd := func(c *mgo.Collection) error {
		fn := c.Remove(bson.M{ "system.id": baseSystem.System.Id})
		return fn
	}

	remove := func() error {
		return withCollection(collection, cmd)
	}

	ierr := remove()
	if ierr != nil {
		// note:if thedoc is already deletedmust return OK
		// if errr.Error() == '...' {
		//    return nil
		// }
		return  errors.New("Database Error: " + ierr.Error())
	}
	return nil
}