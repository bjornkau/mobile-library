package main

//setting up the database for testing

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"testing"
	"time"
)

func setUpDB(t *testing.T) *APIMongoDB {
	db := APIMongoDB{
		"mongodb://test:m0bile@ds141264.mlab.com:41264/mobilelibrary",
		"mobilelibrary",
		"testbooks",
		"testusers",
	}

	session, err := mgo.Dial(db.Host)
	defer session.Close()

	if err != nil {
		t.Error(err)
	}
	return &db
}

//retrieves info on Player's Handbook
func getBook1Info() (Book, bool) {
	temp, _ := RetrieveBookInfo(9780060853983)
	return ParseBookInfo(temp, 9780060853983)
}

//retrives info on Good Omens
func getBook2Info() (Book, bool) {
	temp, _ := RetrieveBookInfo(9780786965601)
	return ParseBookInfo(temp, 9780786965601)
}

func getUserInfo() (user User) {
	user.Username = "temp"
	user.Password = "password"
	user.Email = "user@test.com"
	user.BooksISBN = []int{9780786965601, 9780060853983}
	tempStatus := StatusStruct{}
	tempStatus.Online = true
	tempStatus.TimeStamp = time.Now().Unix()
	user.Status = tempStatus
	return
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

func TestAPIMongoDB_User(t *testing.T) {
	db := setUpDB(t)
	defer tearDownDB(t, db)

	db.Init()
	if db.CountUsers() != 0 {
		t.Error("database not properly initiated, should be empty")
	}

	user := getUserInfo()
	_ = db.AddUser(user)

	if db.CountUsers() != 1 {
		t.Error("Did not insert properly")
	}

	user2, err := db.GetUserByUsername(user.Username)

	if user.Username != user2.Username || user.Password != user2.Password || !err {
		t.Error("Did not find user")
	}
	user.Status.Online = false
	updateErr := db.UpdateUserStatus(user)
	if updateErr != nil {
		fmt.Print(updateErr)
		//t.Error("Did not update right")
		t.Error(updateErr)
	}

	user2, _ = db.GetUserByUsername(user.Username)
	if user2.Status.Online {
		t.Error("Update failed")
	}

	db.DeleteUser(user)

	if db.CountUsers() != 0 {
		t.Error("Did not delete properly")
	}

}
