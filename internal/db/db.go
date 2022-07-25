package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Database interface {
	Connect() (*sql.DB, error)
}

type db struct {
	username string
	password string
	host     string
	port     string
	dbName   string
}

func New(username, password, host, port, dbName string) db {
	return db{
		username,
		password,
		host,
		port,
		dbName,
	}
}

func (d db) Connect() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", d.username, d.password, d.host, d.port, d.dbName),
	)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
