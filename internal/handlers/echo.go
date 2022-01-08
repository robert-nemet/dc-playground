package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dc-playground/model"
	"dc-playground/services"
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
	var body []byte
	_, err := r.Body.Read(body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	json.Unmarshal(body, &msg)

	rsp := e.svc.EchoMsg(msg)

	w.Write([]byte(fmt.Sprintf("{ 'echo': %s}", rsp)))
}
