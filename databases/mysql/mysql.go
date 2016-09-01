package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testIris/config"
)

var session *sql.DB

func Conn() *sql.DB {
	return session
}

func init() {
	db, err := sql.Open("mysql", config.MysqlUrl)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	session = db
}
