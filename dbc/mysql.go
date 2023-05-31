package dbc

import (
	"fmt"
	"gotest/log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var lock = &sync.Mutex{}

var db *sqlx.DB

var mylog = log.NewFlieLogger("debug", "./", "now.log", 100*1024*1024)

func MySQL() *sqlx.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			db = initClient()

		}
	}
	return db
}

func initClient() *sqlx.DB {
	cfg, err := configure("./conf/psql.json")
	if err != nil {
		mylog.Fatal(err.Error())

	}
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DatabaseName)
	database, err := sqlx.Open("mysql", dbUrl)
	if err != nil {
		fmt.Println(dbUrl)
		mylog.Fatal(err.Error())
	}
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
	if err != nil {
		mylog.Fatal(err.Error())
	}
	db = database
	return db
}
