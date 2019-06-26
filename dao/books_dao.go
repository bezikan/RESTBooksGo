package dao

import (
	"log"

	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BooksDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "books"
)

func (m *BooksDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *BooksDAO) FindAll() ([]Book, error) {
	var books []Book
	err := db.C(COLLECTION).Find(bson.M{}).All(&books)
	return books, err
}

func (m *BooksDAO) FindById(id string) (Book, error) {
	var book Book
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

func (m *BooksDAO) Insert(book Book) error {
	err := db.C(COLLECTION).Insert(&book)
	return err
}

func (m *BooksDAO) Delete(book Book) error {
	err := db.C(COLLECTION).Remove(&book)
	return err
}

func (m *BooksDAO) Update(book Book) error {
	err := db.C(COLLECTION).UpdateId(book.ID, &book)
	return err
}
