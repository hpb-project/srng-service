package models

import (
	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
	"log"
)

var globalS *mgo.Session
var db string
var collection string

func init() {
	host := beego.AppConfig.String("mongo::host")
	db = beego.AppConfig.String("mongo::db")
	collection = beego.AppConfig.String("mongo::collection")
	diaInfo := &mgo.DialInfo{
		Addrs: []string{host},
	}
	s, err := mgo.DialWithInfo(diaInfo)

	if err != nil {
		log.Fatalln("create session error", err)
	}

	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := globalS.Copy()
	c := s.DB(db).C(collection)
	return s, c
}

func Insert(docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Insert(docs...)
}

func FindPage(pageNo, pageSize int, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(nil).Skip(pageNo).Limit(pageSize).All(result)
}

func FindAllByHash(query, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).All(result)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func Update(db, collection string, query, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Update(query, update)
}

func Remove(db, collection string, query interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Remove(query)
}
