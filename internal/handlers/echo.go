package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"dc-playground/internal/model"
	"dc-playground/internal/services"
)

type EchoHandlers interface {
	EchoHandler(w http.ResponseWriter, r *http.Request)
}

type echohndl struct {
	svc services.EchoSvc
}

func NewEchoHandler(svc services.EchoSvc) EchoHandlers {
	return echohndl{
		svc: svc,
	}
}

func (e echohndl) EchoHandler(w http.ResponseWriter, r *http.Request) {
	var msg model.Echo
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.EchoRsp{Rsp: err.Error()})
		return
	}

	if rand.Intn(100) < 10 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.EchoRsp{
			Rsp: "Internal Server Error",
		})
		return
	}

	rsp := e.svc.EchoMsg(msg)

	er := model.EchoRsp{
		Rsp: rsp,
	}

	json.NewEncoder(w).Encode(er)
}
