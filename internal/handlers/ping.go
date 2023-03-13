package handlers

import (
	"dc-playground/internal/model"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type PingHandlers interface {
	PingHandler(w http.ResponseWriter, r *http.Request)
}

type pingHandler struct{}

func NewPingHandler() PingHandlers {
	return &pingHandler{}
}

func (p *pingHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	var msg model.Ping
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Pong{Msg: err.Error()})
		return
	}

	if rand.Intn(100) < 10 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Pong{
			Msg: "Internal Server Error",
		})
		return
	}

	received := time.Now().Format(time.RFC3339)
	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	responded := time.Now().Format(time.RFC3339)

	json.NewEncoder(w).Encode(model.Pong{
		Msg:          msg.Msg,
		TimeReceived: received,
		TimeSent:     responded,
	})
}
