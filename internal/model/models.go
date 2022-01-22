package model

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
