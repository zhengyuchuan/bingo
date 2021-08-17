package db

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var MongoSession *mgo.Session

func InitMongo() {
	var err error
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	MongoSession, err = mgo.Dial("mongodb://@localhost:40001?maxPoolSize=100")
	if err != nil {
		fmt.Println("connect mongo failed")
	}
}

func exampleMongo() {
	s := MongoSession.Copy()
	defer s.Close()

	s.DB("db").C("table").Insert(bson.M{"name": "zheng"})
}
