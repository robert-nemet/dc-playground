package handlers

import (
	"dc-playground/internal/model"
	"dc-playground/internal/services"
	"encoding/json"
	"net/http"
)

type CounterHandler interface {
	SaveCount(w http.ResponseWriter, r *http.Request)
	ReadCount(w http.ResponseWriter, r *http.Request)
}

type counterHandler struct {
	svc services.DB
}

func NewCounterHandler(svc services.DB) CounterHandler {
	return counterHandler{
		svc: svc,
	}
}

func (ch counterHandler) SaveCount(w http.ResponseWriter, r *http.Request) {
	err := ch.svc.SaveCnt()
	rsp := model.IncCountRsp{
		Msg: "OK",
	}

	if err != nil {
		rsp.Msg = err.Error()
	}

	json.NewEncoder(w).Encode(rsp)
}

func (ch counterHandler) ReadCount(w http.ResponseWriter, r *http.Request) {
	v, err := ch.svc.GetCnt()
	if err != nil {
		rsp := model.IncCountRsp{
			Msg: err.Error(),
		}
		json.NewEncoder(w).Encode(rsp)
		return
	}
	rsp := model.CountRsp{
		Cnt: v,
	}
	json.NewEncoder(w).Encode(rsp)
}
