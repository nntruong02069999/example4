package database

import (
	"errors"
	"fmt"
	_"log"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Id         string `xorm:"unique" json:"id"`
	Name       string `xorm:"not null" json:"name"`
	Birth      int64	`json:"birth"`
	Created    int64 `xorm:"created" json:"created"`
	UpdatedAt int64  `json:"update_at"`
}

func (database *Db) CreateUser(user *User) error {
	aff , err := database.engine.Insert(user)
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("Cannot insert to table User")
	}
	return nil
}

func (database *Db) GetListUser() (*[]User, error) {
	var user *[]User
	err := database.engine.Find(user)
	if err != nil {
		return nil, errors.New("Can not find list")
	}
	return user , nil
}

func (database *Db) GetUserById(id string) (*User, error) {
	user := &User{Id: id}
	has, err := database.engine.Get(user)
	if err != nil {
		return nil , err
	}
	if !has {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (database *Db) UpdateUser(object, conditions *User) (error) {
	aff , err := database.engine.Update(object, conditions)
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("cannot update user")
	}
	return nil
}

func (database *Db) InsertToPointAfterCreateUser(user *User) (error) {
	err := database.CreateUser(user)
	if err != nil {
		return err
	}
	point := &Point{UserId : user.Id, Points : 10}
	err = database.CreatePoint(point)
	if err != nil {
		return err
	}
	return nil
}


func (database *Db) UpdateBirthUser(birth int64, id string) (error) {
	session := database.engine.NewSession()
	defer session.Close()
	if err := session.Begin() ; err != nil {
		return err
	}
	// Check user exits
	user := &User{Id : id}
	has , err := session.Get(user)
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
	 _ , err = session.Update(user, &User{Id : id})
	 if err != nil {
		 session.Rollback()
		 return err
	 }

	 // Update point of user
	 point := &Point{UserId: user.Id}
	 _ , err2 := session.Get(point)
	 if err2 != nil {
		 session.Rollback()
		 return err2
	 }
	 point.Points += 10
	 _ , err2 = session.Update(point, &Point{UserId: user.Id})
	 if err2 != nil {
		session.Rollback()
		return err2
	 }
	 session.Commit()
	 return nil
}

func (database *Db) TestInsertUserUsingGoroutines() {
	dsUser := NewDsUser()
	for i := 1 ; i <= 10 ; i++ {
		go func() {
			for i := 1 ; i <= 10 ; i++ {
				dsUser.InsertNewUser(database)
			}
		}()
	}
	time.Sleep(5 * time.Second)
}

func (data *DsDataUser) InsertNewUser(database *Db){
	data.Lock()
	var user User
	user.Id = strconv.FormatInt(data.indentity, 10) 
	user.Name = "Test " + user.Id  
	data.indentity++
	database.CreateUser(&user)
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

func (db *Db) ScanTableUser(buffData chan *dataUser, wg *sync.WaitGroup) (error) {
	
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
			dtuser := &dataUser{user: *user, indentity: count}
			count ++
			buffData <- dtuser
			wg.Add(1)
		}
	}
	return nil
}

func Bai3(database *Db) (error) {
	buffData := make(chan *dataUser, 100)
	defer close(buffData)
	var wg sync.WaitGroup
	for i := 1 ; i <=2 ; i++ {
		go printData(buffData, &wg)
	}
	err := database.ScanTableUser(buffData, &wg)
	
	if err != nil {
		return err
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
