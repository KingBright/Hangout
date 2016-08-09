package model

import (
	"log"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
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
