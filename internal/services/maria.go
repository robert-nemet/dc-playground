package services

import (
	"database/sql"
	"dc-playground/internal/config"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	writeSqlStatement  = "INSERT INTO `counters` (`key`, `counter`) VALUES (?, ?);"
	updateSqlStatement = "UPDATE `counters` SET `counter`=? WHERE `key`=?;"
	readSqlStatement   = "SELECT `counter` FROM `counters` WHERE `key`=?;"
)

type maria_database struct {
	db *sql.DB
}

func NewMariaDBService(cfg config.AppConfig) DB {
	var db *sql.DB
	mariaInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	fmt.Println(mariaInfo)
	db, _ = sql.Open("mysql", mariaInfo)
	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	return maria_database{
		db: db,
	}
}

func (db maria_database) GetCnt() (i int, e error) {
	weekday := time.Now().Weekday()

	stmt, err := db.db.Prepare(readSqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var counter int
	err = stmt.QueryRow(weekday).Scan(&counter)
	if err != nil {
		log.Println(err)
		if e == sql.ErrNoRows {
			log.Printf("No rows for %s day.", weekday)
			counter = 0
			e = nil
		}
	}

	log.Println("Read counter ", counter)

	return counter, e
}

func (db maria_database) SaveCnt() error {
	weekday := time.Now().Weekday()
	counter, err := db.GetCnt()
	if err == nil {
		log.Printf("Inc for %s, cnt %d", weekday, counter)
		counter++
		if counter == 1 {
			log.Println("Write new ", weekday, " ", counter)
			_, err = db.db.Exec(writeSqlStatement, weekday, counter)
		} else {
			log.Println("Update new ", weekday, " ", counter)
			_, err = db.db.Exec(updateSqlStatement, counter, weekday)
			if err != nil {
				log.Println("Update error:", err)
			}
		}
	}
	return err
}
