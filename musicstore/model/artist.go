package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Artist struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name"`
	Genres []Genre       `json:"genres"`
	Born   *time.Time    `json:"born"`
}
