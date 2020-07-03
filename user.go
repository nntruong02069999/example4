package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/nntruong02069999/example4/database"
)

func InsertToPointAfterCreateUser(user *database.User) error {
	err := db.CreateUser(user)
	if err != nil {
		return err
	}
	point := &database.Point{UserId: user.Id, Points: 10}
	err = db.CreatePoint(point)
	if err != nil {
		return err
	}
	return nil
}

func TestInsertUserUsingGoroutines() {
	dsUser := NewDsUser()
	for i := 1; i <= 10; i++ {
		go func() {
			for i := 1; i <= 10; i++ {
				 InsertNewUser(dsUser)
			}
		}()
	}
	time.Sleep(5 * time.Second)
}

func InsertNewUser(data *database.DsDataUser) error {
	data.Lock()
	defer data.Unlock()
	var user *database.User
	user.Id = strconv.FormatInt(data.Indentity, 10)
	user.Name = "Test " + user.Id
	data.Indentity++
	err := db.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func NewDsUser() *database.DsDataUser {
	return &database.DsDataUser{
		Indentity: 1,
	}
}

func Bai3() error {
	buffData := make(chan *database.DataUser, 100)
	defer close(buffData)
	var wg sync.WaitGroup
	for i := 1; i <= 2; i++ {
		go printData(buffData, &wg)
	}
	err := db.ScanTableUser(buffData, &wg)

	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func printData(buffData chan *database.DataUser, wg *sync.WaitGroup) {
	for {
		select {
		case data := <-buffData:
			fmt.Printf("counter%v - %v - %v\n", data.Indentity, data.DataUser.Id, data.DataUser.Name)
			wg.Done()
		}
	}
}
