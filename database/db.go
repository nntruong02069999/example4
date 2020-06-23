package database

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"log"

)

type Db struct {
	engine *xorm.Engine
}

var (
	tables [] interface{}
)

func (db *Db) ConnectDb(){
	var err error
	db.engine, err = xorm.NewEngine("mysql","root:root@/test?charset=utf8")
	if err != nil {
		log.Println("Connect database faild")
	} else {
		log.Println("Connect database success")
	}
	db.engine.ShowSQL(true)
}

func (db *Db) InitDatabase(){
	initTables()
	db.engine.Sync2(tables...)
}

func initTables(){
	tables = append(tables, new(User), new(Point))
}