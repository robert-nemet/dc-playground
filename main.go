package main

import (
	"log"
	"net/http"

	"dc-playground/internal/handlers"
	"dc-playground/internal/services"

	"github.com/gorilla/mux"
)

func main() {

	es := services.NewEchoSvc()
	eh := handlers.NewEchoHandler(es)

	r := mux.NewRouter()
	r.HandleFunc("/echo", eh.EchoHandler).Methods(http.MethodPost)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9999", nil))

}
