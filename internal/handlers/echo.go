package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"dc-playground/internal/config"
	"dc-playground/internal/model"
)

type EchoHandlers interface {
	EchoHandler(w http.ResponseWriter, r *http.Request)
}

type echohndl struct {
	cfg config.AppConfig
}

func NewEchoHandler(cfg config.AppConfig) EchoHandlers {
	return echohndl{
		cfg: cfg,
	}
}

func (e echohndl) EchoHandler(w http.ResponseWriter, r *http.Request) {
	var msg model.Echo
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.EchoRsp{Rsp: err.Error()})
		return
	}

	if rand.Intn(100) < e.cfg.ErrorRate {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.EchoRsp{
			Rsp: "Internal Server Error",
		})
		return
	}

	rsp := msg.Msg

	er := model.EchoRsp{
		Rsp: rsp,
	}

	json.NewEncoder(w).Encode(er)
}
