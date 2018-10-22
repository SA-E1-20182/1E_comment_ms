package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type (
	Comment struct {
		Id         	bson.ObjectId 	`bson:"_id,omitempty" json:"id"`
		version_id  int	        	`json:"version_id"`
		image_id  	int		       	`json:"image_id"`
		message		string        	`json:"message"`
		likes		int	        	`json:"likes"`
		created_by  string     		`json:"created_by"`
		created_at  time.Time     	`json:"created_at,omitempty"`
	}
)
