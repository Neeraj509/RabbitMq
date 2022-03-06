package routers

import (
	C "example.com/m/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", C.Input).Methods("POST")

	return router
}
