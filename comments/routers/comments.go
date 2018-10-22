package routers

import (
	"github.com/gorilla/mux"
	"krajono/comments/controllers"
	"log"
)

func SetCommentsRouters(router *mux.Router) *mux.Router {
	log.Printf("Set comments routers")
	router.HandleFunc("/comments", controllers.GetComments).Methods("GET")
	router.HandleFunc("/comments", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/comments/{id}", controllers.GetCommentById).Methods("GET")
	router.HandleFunc("/comments/{id}", controllers.DeleteComment).Methods("DELETE")
	return router
}
