package controllers

import (
	"encoding/json"
	"net/http"
	"krajono/comments/common"
	"krajono/comments/data"
)

// Handler for HTTP Post - "/comments"
// Create a new Comment document
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var dataResource CommentResource
	// Decode the incoming Comment json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Comment data", 500)
		return
	}
	comment := &dataResource.Data
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("comments")
	// Create Comment
	repo := &data.CommentRepository{c}
	repo.Create(comment)
	// Create response data
	j, err := json.Marshal(dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetComments(w http.ResponseWriter, r *http.Request) {
	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("comments")
	repo := &data.CommentRepository{c}
	// Get all comments
	comments := repo.GetAll()
	// Create response data
	j, err := json.Marshal(CommentsResource{Data: comments})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
