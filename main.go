package main

import (
	"fmt"
	"log"
	"net/http"

	router "example.com/m/routers"
)

func main() {
	log.Println("Starting the application")

	r := router.Router()

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")

}
