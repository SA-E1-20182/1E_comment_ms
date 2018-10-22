package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type (
	Comment struct {
		Id         	bson.ObjectId 	`bson:"_id,omitempty" json:"id"`
		Version_id  int	        	`json:"version_id"`
		Image_id  	int		       	`json:"image_id"`
		Message		string        	`json:"message"`
		Likes		int	        	`json:"likes"`
		Created_by  string     		`json:"created_by"`
		Created_at  time.Time     	`json:"created_at,omitempty"`
	}
)
