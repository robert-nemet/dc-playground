package main

import (
	"log"
	"net/http"

	"dc-playground/handlers"
	"dc-playground/services"
)

func main() {

	es := services.NewEchoSvc()
	eh := handlers.NewEchoHandler(es)

	http.HandleFunc("/echo", eh.EchoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
