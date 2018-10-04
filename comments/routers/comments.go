package routers

import (
	"github.com/gorilla/mux"
	"../controllers"
)

func SetBookingsRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/comments", controllers.GetBookings).Methods("GET")
	router.HandleFunc("/comments", controllers.CreateBooking).Methods("POST")
	return router
}
