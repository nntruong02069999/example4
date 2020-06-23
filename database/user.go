package database

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Id         string `xorm:"unique"`
	Name       string `xorm:"not null"`
	Birth      int64
	Created    int64 `xorm:"created"`
	Updated_at int64 `xorm:"updated"`
}

func (db *Db) CreateUser(user *User) error {
	_ , err := db.engine.Insert(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db *Db) GetListUser() ([]User, error) {
	var user []User
	err := db.engine.Find(&user)
	if err != nil {
		log.Println("Không tìm thấy danh sách user")
		return nil, err
	}
	return user , nil
}

func (db *Db) GetUserById(id string) (*User, error) {
	user := &User{Id: id}
	has, err := db.engine.Get(user)
	if err != nil {
		log.Println("Failed")
		return nil , err
	}
	if !has {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (db *Db) UpdateUser(user *User) (error) {
	_, err := db.engine.Update(user)
	if err != nil {
		log.Println("Update failed")
		return err
	}
	return nil
}

func (db *Db) InsertToPointAfterCreateUser(user *User) (error) {
	err := db.CreateUser(user)
	if err != nil {
		log.Println(err)
		return err
	}
	point := &Point{User_id : user.Id, Points : 10}
	err = db.CreatePoint(point)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}


func (db *Db) UpdateBirthUser(birth int64, id string) (error) {
	session := db.engine.NewSession()
	defer session.Close()
	if err := session.Begin() ; err != nil {
		log.Println(err)
		return err
	}
	// Check user exits
	 user , err := db.GetUserById(id)
	 if err != nil {
		 log.Println(err)
		 session.Rollback()
		 return err
	 }
	 // Update info user
	 user.Birth = birth
	 user.Name = user.Name + " updated"
	 user.Updated_at = time.Now().UnixNano()
	 err = db.UpdateUser(user)
	 if err != nil {
		 log.Println(err)
		 session.Rollback()
		 return err
	 }

	 // Update point of user
	 point, err2 := db.GetPointById(user.Id)
	 if err2 != nil {
		 log.Println(err2)
		 session.Rollback()
		 return err2
	 }
	 point.Points += 10
	 err2 = db.UpdatePoint(point)
	 if err2 != nil {
		log.Println(err2)
		session.Rollback()
		return err2
	 }
	 session.Commit()
	 return nil
}

func (db *Db) TestInsertUserUsingGoroutines() {
	dsUser := NewDsUser()
	for i := 1 ; i <= 10 ; i++ {
		go func() {
			for i := 1 ; i <= 10 ; i++ {
				dsUser.InsertNewUser(db)
			}
		}()
	}
	time.Sleep(5 * time.Second)
}

func (data *DsDataUser) InsertNewUser(db *Db){
	data.Lock()
	var user User
	user.Id = strconv.FormatInt(data.indentity, 10) 
	user.Name = "Test " + user.Id  
	data.indentity++
	db.CreateUser(&user)
	defer data.Unlock()
}

type DsDataUser struct {
	sync.Mutex
	indentity int64
}

func NewDsUser() *DsDataUser {
	return &DsDataUser{
		indentity: 1,
	}
}

type dataUser struct {
	user User
	indentity int
}

func (db *Db) Bai3() (error) {
	buffData := make(chan *dataUser, 100)
	defer close(buffData)
	var wg sync.WaitGroup
	for i := 1 ; i <=2 ; i++ {
		go printData(buffData, &wg)
	}
	rows, err := db.engine.Rows(&User{})
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	user := new(User)
	count := 1
	for rows.Next() {
		rows.Scan(user)
		dtuser := &dataUser{user: *user, indentity: count}
		count ++
		buffData <- dtuser
		wg.Add(1)
	}
	wg.Wait()
	return nil
}

func printData(buffData chan *dataUser, wg *sync.WaitGroup) {
	for {
		select {
		case data := <- buffData :
			fmt.Printf("counter%v - %v - %v\n",data.indentity, data.user.Id, data.user.Name)
			wg.Done()
		}
	}
}
