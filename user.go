package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	_"time"

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

func TestInsertUser() {
	var user database.User
	for i := 1; i <= 100; i++ {
		user.Id = strconv.FormatInt(int64(i), 10)
		user.Name = "Test " + user.Id
		err := db.CreateUser(&user)
		if err != nil {
			log.Println(err)
		}
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
