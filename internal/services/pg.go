package services

import (
	"database/sql"
	"dc-playground/internal/config"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	writeSqlStatement  = `INSERT INTO counters (key, counter) VALUES ($1, $2)`
	updateSqlStatement = `UPDATE counters SET counter=$2 WHERE key=$1`
	readSqlStatement   = `SELECT counter FROM counters where key=$1`
)

type DB interface {
	GetCnt() (i int, e error)
	SaveCnt() error
}

type database struct {
	db *sql.DB
}

func NewDBService(cfg config.AppConfig) DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	fmt.Println(psqlInfo)
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

func (db database) GetCnt() (i int, e error) {
	weekday := time.Now().Weekday()
	var cnt int
	row := db.db.QueryRow(readSqlStatement, weekday)
	if e = row.Scan(&cnt); e == sql.ErrNoRows {
		log.Printf("No rows for %s day.", weekday)
		cnt = 0
		e = nil
	}
	return cnt, e
}

func (db database) SaveCnt() error {
	weekday := time.Now().Weekday()
	counter, err := db.GetCnt()
	if err == nil {
		log.Printf("Inc for %s, cnt %d", weekday, counter)
		counter++
		if counter == 1 {
			_, err = db.db.Exec(writeSqlStatement, weekday, counter)
		} else {
			_, err = db.db.Exec(updateSqlStatement, weekday, counter)
		}
	}
	return err
}
