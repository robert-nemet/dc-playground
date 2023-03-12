package main

import (
	"fmt"
	"log"
	"net/http"

	"dc-playground/internal/config"
	"dc-playground/internal/handlers"
	"dc-playground/internal/middleware"
	"dc-playground/internal/services"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	r.HandleFunc("/echo", middleware.Instrument(eh.EchoHandler)).Methods(http.MethodPost)
	r.HandleFunc("/counter", middleware.Instrument(ch.SaveCount)).Methods(http.MethodPost)
	r.HandleFunc("/counter", middleware.Instrument(ch.ReadCount)).Methods(http.MethodGet)

	r.HandleFunc("/ping", middleware.Instrument(handlers.NewPingHandler().PingHandler)).Methods(http.MethodPost)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	port := fmt.Sprintf(":%v", cfg.AppPort)
	fmt.Println("Start on " + port)
	log.Fatal(http.ListenAndServe(port, nil))

}
