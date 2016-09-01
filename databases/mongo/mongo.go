package mongo

import (
	"gopkg.in/mgo.v2"
	"testIris/config"
)

var session *mgo.Session

func Conn() *mgo.Session {
	return session.Copy()
}

func init() {
	session, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Primary, true)
}
