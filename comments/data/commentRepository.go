package data

import (
	"krajono/comments/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CommentRepository struct {
	C *mgo.Collection
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	obj_id := bson.NewObjectId()
	comment.Id = obj_id
	err := r.C.Insert(&comment)
	return err
}

func (r *CommentRepository) GetAll() []models.Comment {
	var comments [] models.Comment
	iter := r.C.Find(nil).Iter()
	result := models.Comment{}
	for iter.Next(&result) {
		comments = append(comments, result)
	}
	return comments
}

func (r *CommentRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
