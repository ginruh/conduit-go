package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type DB struct {
	Engine *xorm.Engine
}

func (db DB) Connect(username, password, host, port, dbName string) error {
	engine, err := xorm.NewEngine(
		"mysql",
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, dbName),
	)
	if err != nil {
		return err
	}
	db.Engine = engine
	return nil
}
