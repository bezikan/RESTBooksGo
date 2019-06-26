package models

import "gopkg.in/mgo.v2/bson"

type Book struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Title  string        `bson:"title" json:"title"`
	Author string        `bson:"author" json:"suthor"`
	ISBN   string        `bson:"isbn" json:"jsbn"`
}
