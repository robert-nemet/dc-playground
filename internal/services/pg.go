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
	pg_writeSqlStatement  = `INSERT INTO counters (key, counter) VALUES ($1, $2)`
	pg_updateSqlStatement = `UPDATE counters SET counter=$2 WHERE key=$1`
	pg_readSqlStatement   = `SELECT counter FROM counters WHERE key=$1`
)

type pgdatabase struct {
	db *sql.DB
}

func NewPGService(cfg config.AppConfig) DB {
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
	return pgdatabase{
		db: db,
	}
}

func (db pgdatabase) GetCnt() (i int, e error) {
	weekday := time.Now().Weekday()
	var cnt int
	row := db.db.QueryRow(pg_readSqlStatement, weekday)
	if e = row.Scan(&cnt); e == sql.ErrNoRows {
		log.Printf("No rows for %s day.", weekday)
		cnt = 0
		e = nil
	}
	return cnt, e
}

func (db pgdatabase) SaveCnt() error {
	weekday := time.Now().Weekday()
	counter, err := db.GetCnt()
	if err == nil {
		log.Printf("Inc for %s, cnt %d", weekday, counter)
		counter++
		if counter == 1 {
			_, err = db.db.Exec(pg_writeSqlStatement, weekday, counter)
		} else {
			_, err = db.db.Exec(pg_updateSqlStatement, weekday, counter)
		}
	}
	return err
}
