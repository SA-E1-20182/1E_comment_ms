package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Comment struct {
		id         	bson.ObjectId 	`bson:"_id,omitempty" json:"id"`
		author_id  	string        	`json:"author_id"`
		image_id  	string        	`json:"image_id"`
		text		string        	`json:"text"`
		created_at  string     		`json:"created_at"`
	}
)
