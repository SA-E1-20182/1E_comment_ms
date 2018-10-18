package main

import (
	"log"
	"net/http"
	"krajono/comments/common"
	"krajono/comments/routers"
)

// Entry point for the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	log.Printf("end init routes")
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
	if err := http.ListenAndServe(":4100", router); err != nil {
		log.Fatal(err)
	}
}
