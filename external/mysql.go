//go:build mysql
// +build mysql

package external

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var MySQL *sql.DB

func init() {
	var dsn string
	// TODO::: get dsn from OS flag

	var err error
	MySQL, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("[External Dependency] Failed to initialize mysql database. Error: " + err.Error())
	}

	err = MySQL.Ping()
	if err != nil {
		panic("[External Dependency] Failed to connect to mysql database. Error: " + err.Error())
	}
}
