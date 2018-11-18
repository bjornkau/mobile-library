package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Struct definition and init function

//APIMongoDB struct stores methods and variables required for database interaction
type APIMongoDB struct {
	Host               string
	DatabaseName       string
	BookCollectionName string
	UserCollectionName string
}

//Init checks conection to db
func (db *APIMongoDB) Init() {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
}

//DB functions related to books

//AddBook takes param book and adds it to database. Returns error nil when all goes well
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

//CountBooks returns integer with number of books in db, returns -1 when an error occurs
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

//GetBookByISBN takes param isbn. Returns true if found book, and a Book struct containing found info
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

//DeleteBook takes param Book. Deletes entry in db. Returns status bool
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

//DB functions related to users

//AddUser takes param User. Adds to db. Returns nil when all is good
func (db *APIMongoDB) AddUser(u User) error {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	errInsert := session.DB(db.DatabaseName).C(db.UserCollectionName).Insert(u)
	if errInsert != nil {
		fmt.Printf("Error in Insert(): %v", errInsert.Error())
		return errInsert
	}
	return nil
}

//CountUsers returns integer with number of books in db, returns -1 when an error occurs
func (db *APIMongoDB) CountUsers() int {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	count, errCount := session.DB(db.DatabaseName).C(db.UserCollectionName).Count()
	if errCount != nil {
		fmt.Printf("Error in Count(): %v", errCount.Error())
		return -1
	}
	return count
}

//DeleteUser takes param User and removes it from db. Returns status bool
func (db *APIMongoDB) DeleteUser(u User) (allIsWell bool) {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	allIsWell = true
	err2 := session.DB(db.DatabaseName).C(db.UserCollectionName).Remove(u)

	if err2 != nil {
		allIsWell = false
	}
	return
}

//GetUserByUsername takes param username. returns bool and user from db
func (db *APIMongoDB) GetUserByUsername(username string) (User, bool) {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	user := User{}
	found := true
	errFind := session.DB(db.DatabaseName).C(db.UserCollectionName).Find(bson.M{"username": username}).One(&user)
	if errFind != nil {
		found = false
	}
	return user, found
}

//UpdateUserStatus takes param user, updates document in db. returns error
func (db *APIMongoDB) UpdateUserStatus(user User) error {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	err = session.DB(db.DatabaseName).C(db.UserCollectionName).Update(bson.M{"username": user.Username}, user)
	return err
}
