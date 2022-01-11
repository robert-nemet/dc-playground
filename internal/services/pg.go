package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB interface {
	GetCnt() int
}

type database struct {
	db *sql.DB
}

func newDB(host string, port int, user, password, dbname string) DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return database{
		db: db,
	}
}

func (db database) GetCnt() int {

	return 0
}
