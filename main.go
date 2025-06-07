package main

import (
	"fmt"
	"log"
	"mongoapi/routes"
	"net/http"
)

func main() {
	fmt.Println("Lets learn about go with mongo db")
	r := routes.Route()

	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}
