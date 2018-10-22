package controllers

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"krajono/comments/common"
	"krajono/comments/data"
	"krajono/comments/models"
	"gopkg.in/mgo.v2"
)

// Handler for HTTP Post - "/comments"
// Create a new Comment document
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var dataResource models.Comment
	// Decode the incoming Comment json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Comment data", 500)
		return
	}
	comment := &dataResource
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
	log.Println(j)
	respondWithJson(w, http.StatusCreated, comment)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(j)
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
	j, err := json.Marshal(comments)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	log.Println(j)
	respondWithJson(w, http.StatusCreated, comments)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(j)
}
// Handler for HTTP Get - "/comments/{id}"
// Get comment by id
func GetCommentById(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("comments")
	repo := &data.CommentRepository{c}

	// Get comment by id
	comment, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
	}

	j, err := json.Marshal(CommentResource{Data: comment})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Handler for HTTP Delete - "/comments/{id}"
// Delete comment by id
func DeleteComment(rw http.ResponseWriter, req *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(req)
	id := vars["id"]

	// Create new context
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("comments")
	repo := &data.CommentRepository{c}

	err := repo.Delete(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else {
			common.DisplayAppError(rw, err, "An unexpected error has occurred", 500)
			return
		}
	}
	rw.WriteHeader(http.StatusNoContent)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}