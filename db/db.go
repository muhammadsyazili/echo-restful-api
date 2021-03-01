package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muhammadsyazili/echo-rest/config"
)

var db *sql.DB
var err error

func Init()  {
	conf := config.GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("Connection String Error!")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	
	err = db.Ping()
	if err != nil {
		panic("DSN Invalid!")
	}
}

func CreateConn() *sql.DB {
	return db
}