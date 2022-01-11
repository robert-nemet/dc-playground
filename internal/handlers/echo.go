package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatalln(err)
		fmt.Fprint(w, err)
		return
	}

	log.Printf("Received msg: %v\n", msg)

	rsp := e.svc.EchoMsg(msg)

	er := model.EchoRsp{
		Rsp: rsp,
	}

	json.NewEncoder(w).Encode(er)
}
