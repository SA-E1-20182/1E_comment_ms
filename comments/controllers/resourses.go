package controllers

import (
	"krajono/comments/models"
)

type (
	// For Get - /comments
	CommentsResource struct {
		Data [] models.Comment `json:"data"`
	}
	// For Post/Put - /comments
	CommentResource struct {
		Data models.Comment `json:"data"`
	}
)
