package main

import (
	_"fmt"
	_"log"

	"github.com/nntruong02069999/example4/database"
)

func main() {
	db := new(database.Db)
	//db.InitDatabase()
	db.ConnectDb()

	//Create new User
	// var user database.User
	// user.Id = "123456"
	// user.Name = "Nguyễn Nam Trường"
	// db.InsertToPointAfterCreateUser(&user)

	// Get list user
	// users := db.GetListUser()
	// if len(users) > 0 {
	// 	fmt.Println(users)
	// } else {
	// 	log.Println("Không tìm thấy danh sách user")
	// }

	// Get user by ID
	// user := db.GetUserById("123456")
	// if user == (database.User{}){
	// 	log.Println("Không tìm thấy user")
	// } else {
	// 	fmt.Println(user)
	// }

	// Update birth
	//db.UpdateBirthUser(1234567,"123456")

	//Insert 100 user
	//db.TestInsertUserUsingGoroutines()
	// Print data

	database.Bai3(db)
}