package main

import (
	"log"
	"net/http"

	"dc-playground/internal/handlers"
	"dc-playground/internal/services"
)

func main() {

	es := services.NewEchoSvc()
	eh := handlers.NewEchoHandler(es)

	http.HandleFunc("/echo", eh.EchoHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))

}
