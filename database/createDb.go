package database

import (
	"sync"
)

type Point struct {
	UserId    string `json:"user_id"`
	Points    int64  `json:"points"`
	MaxPoints int64  `json:"max_points"`
}

type User struct {
	Id         string `xorm:"unique" json:"id"`
	Name       string `xorm:"not null" json:"name"`
	Birth      int64	`json:"birth"`
	Created    int64 `xorm:"created" json:"created"`
	UpdatedAt int64  `json:"updated_at"`
}


//// ------------ For testing ------------


type DsDataUser struct {
	sync.Mutex
	Indentity int64
}

type DataUser struct {
	DataUser User
	Indentity int
}