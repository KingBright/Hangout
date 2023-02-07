package model

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/core"
)

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./hangout-test.db")
	if err != nil {
		log.Fatalf("failed init with %s", err.Error())
	}
	engine.SetMapper(core.GonicMapper{})
	initDB()
}
