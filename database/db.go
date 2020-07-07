package database

import (
	"errors"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Db struct {
	engine *xorm.Engine
}

var (
	tables []interface{}
)

func (db *Db) ConnectDb() error {
	var err error
	db.engine, err = xorm.NewEngine("mysql", "truong:root@/test?charset=utf8")
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

// ---------------  User----------------------

func (database *Db) CreateUser(user *User) error {
	aff, err := database.engine.Insert(user)
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("Cannot insert to table User")
	}
	return nil
}

func (database *Db) GetListUsers() ([]*User, error) {
	var user []*User
	err := database.engine.Find(user)
	if err != nil {
		return nil, errors.New("Can not find list")
	}
	return user, nil
}

func (database *Db) GetUserById(id string) (*User, error) {
	user := &User{Id: id}
	has, err := database.engine.Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (database *Db) UpdateUser(object, conditions *User) error {
	aff, err := database.engine.Update(object, conditions)
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("cannot update user")
	}
	return nil
}

func (database *Db) UpdateBirthUser(birth int64, id string) error {
	session := database.engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	// Check user exits
	user := &User{Id: id}
	has, err := session.Get(user)
	if err != nil {
		session.Rollback()
		return err
	}
	if !has {
		session.Rollback()
		return errors.New("User not found")
	}
	// Update info user
	user.Birth = birth
	user.Name = user.Name + " updated"
	user.UpdatedAt = time.Now().UnixNano()
	_, err = session.Update(user, &User{Id: id})
	if err != nil {
		session.Rollback()
		return err
	}

	// Update point of user
	point := &Point{UserId: user.Id}
	_, err2 := session.Get(point)
	if err2 != nil {
		session.Rollback()
		return err2
	}
	point.Points += 10
	_, err2 = session.Update(point, &Point{UserId: user.Id})
	if err2 != nil {
		session.Rollback()
		return err2
	}
	session.Commit()
	return nil
}

func (db *Db) ScanTableUser(buffData chan *DataUser, wg *sync.WaitGroup) error {

	rows, err := db.engine.Rows(&User{})
	if err != nil {
		return err
	}
	defer rows.Close()
	user := new(User)
	count := 1
	for rows.Next() {
		err2 := rows.Scan(user)
		if err2 == nil {
			dtuser := &DataUser{DataUser: *user, Indentity: count}
			count++
			buffData <- dtuser
			wg.Add(1)
		}
	}
	return nil
}

// ------------------- Point -------------------------------

func (database *Db) CreatePoint(point *Point) error {
	aff, err := database.engine.Insert(point)
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("Cannot insert")
	}
	return nil
}

func (database *Db) UpdatePoint(point, conditions *Point) error {
	aff, err := database.engine.Update(point, conditions)
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("Cannot update point")
	}
	return nil
}

func (database *Db) GetPointById(id string) (*Point, error) {
	point := &Point{UserId: id}
	has, err := database.engine.Get(point)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("not found")
	}
	return point, nil
}
