package controllers

import (
	"../models"
)

type (
	// For Get - /comments
	CommentsResource struct {
		Data []models.Comments `json:"data"`
	}
	// For Post/Put - /comments
	CommentsResource struct {
		Data models.Comments `json:"data"`
	}
)
