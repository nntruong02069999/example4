package database

import (
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Db struct {
	engine *xorm.Engine
}

var (
	tables [] interface{}
)

func (db *Db) ConnectDb() error{
	var err error
	db.engine, err = xorm.NewEngine("mysql","root:root@/test?charset=utf8")
	if err != nil {
		return errors.New("Connect database faild")
	}
	log.Println("Connect database success")
	db.engine.ShowSQL(true)
	return nil
}

func (db *Db) InitDatabase() error {
	initTables()
	err := db.engine.CreateTables(tables...)
	if err != nil {
		return err
	}
	return nil
}

func initTables() {
	tables = append(tables, new(User), new(Point))
}