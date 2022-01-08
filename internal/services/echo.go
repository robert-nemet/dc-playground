package services

import "dc-playground/internal/model"

type EchoSvc interface {
	EchoMsg(echo model.Echo) string
}

type echosvcimpl struct{}

func NewEchoSvc() EchoSvc {
	return echosvcimpl{}
}

func (e echosvcimpl) EchoMsg(echo model.Echo) string {
	return echo.Msg
}
