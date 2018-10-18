package routers

import (
	"github.com/gorilla/mux"
	"log"
)

func InitRoutes() *mux.Router {
	log.Printf("Init rotes")
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the Comment entity
	router = SetCommentsRouters(router)
	return router
}
