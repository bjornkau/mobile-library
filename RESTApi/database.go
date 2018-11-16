package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type APIMongoDB struct {
	Host               string
	DatabaseName       string
	BookCollectionName string
	UserCollectionName string
}

func (db *APIMongoDB) Init() {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
}

func (db *APIMongoDB) AddBook(b Book) error {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	errInsert := session.DB(db.DatabaseName).C(db.BookCollectionName).Insert(b)
	if errInsert != nil {
		fmt.Printf("Error in Insert(): %v", errInsert.Error())
		return errInsert
	}
	return nil
}

func (db *APIMongoDB) CountBooks() int {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	count, errCount := session.DB(db.DatabaseName).C(db.BookCollectionName).Count()
	if errCount != nil {
		fmt.Printf("Error in Count(): %v", errCount.Error())
		return -1
	}
	return count
}

func (db *APIMongoDB) GetBookByISBN(isbn int) (Book, bool) {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	book := Book{}
	found := true
	errFind := session.DB(db.DatabaseName).C(db.BookCollectionName).Find(bson.M{"isbn": isbn}).One(&book)
	if errFind != nil {
		found = false
	}
	return book, found
}

func (db *APIMongoDB) DeleteBook(b Book) (allIsWell bool) {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	allIsWell = true
	err2 := session.DB(db.DatabaseName).C(db.BookCollectionName).Remove(b)

	if err2 != nil {
		allIsWell = false
	}
	return
}
