package model

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request)

type Echo struct {
	Msg string `json:"msg,omitempty"`
}

type EchoRsp struct {
	Rsp string `json:"echo,omitempty"`
}

type CountRsp struct {
	Cnt int `json:"counter"`
}

type IncCountRsp struct {
	Msg string `json:"msg"`
}

type Ping struct {
	Msg      string `json:"msg"`
	TimeSent string `json:"time_sent,omitempty"`
}

type Pong struct {
	Msg          string `json:"msg,omitempty"`
	TimeReceived string `json:"time_received,omitempty"`
	TimeSent     string `json:"time_sent,omitempty"`
}
