package main

import (
	"fmt"
	"log"
	"net/http"

	"dc-playground/internal/config"
	"dc-playground/internal/handlers"
	"dc-playground/internal/services"

	"github.com/gorilla/mux"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	cs := services.NewDBService(cfg)
	ch := handlers.NewCounterHandler(cs)

	es := services.NewEchoSvc()
	eh := handlers.NewEchoHandler(es)

	r := mux.NewRouter()
	r.HandleFunc("/echo", eh.EchoHandler).Methods(http.MethodPost)
	r.HandleFunc("/counter", ch.SaveCount).Methods(http.MethodPost)
	r.HandleFunc("/counter", ch.ReadCount).Methods(http.MethodGet)

	http.Handle("/", r)
	port := fmt.Sprintf(":%v", cfg.AppPort)
	fmt.Println("Start on " + port)
	log.Fatal(http.ListenAndServe(port, nil))

}
