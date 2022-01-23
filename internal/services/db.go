package services

import (
	"dc-playground/internal/config"
)

type DB interface {
	GetCnt() (i int, e error)
	SaveCnt() error
}

func NewDBService(cfg config.AppConfig) DB {
	if cfg.DbType != "PG" {
		return NewMariaDBService(cfg)
	}
	return NewPGService(cfg)
}
