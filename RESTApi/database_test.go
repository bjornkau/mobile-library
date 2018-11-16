package main

//setting up the database for testing

import (
	"gopkg.in/mgo.v2"
	"testing"
)

func setUpDB(t *testing.T) *APIMongoDB {
	db := APIMongoDB{
		"mongodb://127.0.0.1:27017",
		"testParagliding",
		"books",
		"users",
	}

	session, err := mgo.Dial(db.Host)
	defer session.Close()

	if err != nil {
		t.Error(err)
	}
	return &db
}

//retrieves info on Mort
func getBook1Info() (Book, bool) {
	temp, _ := RetrieveBookInfo(9780060853983)
	return ParseBookInfo(temp, 9780060853983)
}

//retrives info on Good Omens
func getBook2Info() (Book, bool) {
	temp, _ := RetrieveBookInfo(9780786965601)
	return ParseBookInfo(temp, 9780786965601)
}

//deleting the database after testing
func tearDownDB(t *testing.T, db *APIMongoDB) {
	session, err := mgo.Dial(db.Host)
	if err != nil {
		t.Error(err)
	}

	err = session.DB(db.DatabaseName).DropDatabase()
	if err != nil {
		t.Error(err)
	}
}

func TestAPIMongoDB_AddBook(t *testing.T) {
	db := setUpDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.CountBooks() != 0 {
		t.Error("database not properly initiated, should be empty")
	}

	book, _ := getBook2Info()
	_ = db.AddBook(book)

	if db.CountBooks() != 1 {
		t.Error("Did not insert properly")
	}
}

func TestAPIMongoDB_GetBook(t *testing.T) {
	db := setUpDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.CountBooks() != 0 {
		t.Error("database not properly initiated, should be empty")
	}

	book, _ := getBook2Info()
	_ = db.AddBook(book)

	if db.CountBooks() != 1 {
		t.Error("Did not insert properly")
	}

	book2, found := db.GetBookByISBN(book.ISBN)
	book3, _ := getBook1Info()

	if !found {
		t.Error("did not find book")
	}

	if book.ISBN != book2.ISBN || book.ISBN == book3.ISBN {
		t.Error("Not equal books")
	}

}

func TestAPIMongoDB_DeleteTrack(t *testing.T) {
	db := setUpDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.CountBooks() != 0 {
		t.Error("database not properly initiated, should be empty")
	}

	book, _ := getBook2Info()
	_ = db.AddBook(book)

	if db.CountBooks() != 1 {
		t.Error("Did not insert properly")
	}

	db.DeleteBook(book)

	if db.CountBooks() != 0 {
		t.Error("Did delete properly")
	}

}
