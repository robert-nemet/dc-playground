package handlers

import (
	"dc-playground/internal/config"
	"dc-playground/internal/model"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type PingHandlers interface {
	PingHandler(w http.ResponseWriter, r *http.Request)
}

type pingHandler struct {
	cfg config.AppConfig
}

func NewPingHandler(cfg config.AppConfig) PingHandlers {
	return &pingHandler{
		cfg: cfg,
	}
}

func (p *pingHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	var msg model.Ping
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Pong{Msg: err.Error()})
		return
	}

	if rand.Intn(100) < p.cfg.ErrorRate {
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
