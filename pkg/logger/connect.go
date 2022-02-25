package logger

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type logger struct {
	conn *sqlx.DB
}

func newconnect() error {
	conn, err := sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true")
	if err != nil {
		fmt.Println("Can't connect")
	}

}

func (l *logger) createTable() {

	_, err := l.conn.Exec(sql)
	if err != nil {
		fmt.Println("Error create table")
	}
}

func (l *logger) InfoError() {
	sql := `INSERT INTO logger() VALUE ()`
}

func (l *logger) InfoSQL() {
	sql := `INSERT INTO logger() VALUE ()`
}

func (l *logger) InfoLog() {
	sql := `INSERT INTO logger() VALUE ()`
}
