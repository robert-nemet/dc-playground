package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"dc-playground/internal/config"
	"dc-playground/internal/handlers"
	"dc-playground/internal/middleware"
	"dc-playground/internal/services"

	"github.com/go-chi/chi/v5"
	chimid "github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	cs := services.NewDBService(cfg)
	ch := handlers.NewCounterHandler(cs)

	eh := handlers.NewEchoHandler(cfg)

	r := chi.NewRouter()
	r.Use(middleware.NewMiddleware().Instrument)
	r.Use(chimid.Logger)
	r.Post("/echo", eh.EchoHandler)
	r.Post("/counter", ch.SaveCount)
	r.Get("/counter", ch.ReadCount)

	r.Post("/ping", handlers.NewPingHandler(cfg).PingHandler)

	r.Handle("/metrics", promhttp.Handler())
	port := fmt.Sprintf(":%v", cfg.AppPort)
	fmt.Println("Start on " + port)

	go func() {
		for {
			body := []byte(`{"msg":"ping"}`)
			_, err := http.Post(cfg.Target+"/ping", "application/json", bytes.NewBuffer(body))
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		}
	}()

	log.Fatal(http.ListenAndServe(port, r))

}
