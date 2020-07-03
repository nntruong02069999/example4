package main

import (
	_ "fmt"
	_ "log"
	"os"
	"github.com/urfave/cli/v2"
	"github.com/nntruong02069999/example4/database"
)

var db = new(database.Db)

func main() {
	err := createCliGolang()
	if err != nil {
		panic("Stop program !!")
	}
	
	TestInsertUserUsingGoroutines()
	//db.ConnectDb()
	//db.InitDatabase()

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

	//Get user by ID
	// user , _ := db.GetUserById("123456")
	// if user == (&database.User{}){
	// 	log.Println("Không tìm thấy user")
	// } else {
	// 	fmt.Println(user)
	// }

	// Update user
	// user := &database.User{}
	// conditionUser := &database.User{Id: "123456"}
	// user.Name = "Testing"
	// user.Birth = 123456666
	// err := db.UpdateUser(user,conditionUser)
	// if err != nil {
	// 	log.Println(err)
	// }
	
	//Update birth
	// err := db.UpdateBirthUser(12345672,"1234562")
	// if err != nil {
	// 	log.Println(err)
	// }	
	//Insert 100 user
	//db.TestInsertUserUsingGoroutines()
	// Print data

	//database.Bai3(db)
}

func createDb(c *cli.Context) error {
	err := db.ConnectDb()
	if err != nil {
		panic(err)
	}
	err = db.InitDatabase()
	if err != nil {
		panic("Create database faild")
	}
	return nil
}

func startApp(c *cli.Context) error {
	err := db.ConnectDb()
	if err != nil {
		panic(err)
	}
	return nil
}

func createCliGolang() error {
	app := cli.NewApp()
	app.Name = "cli-golang"
	app.Version = "0.0.1"
	app.Usage = "Using cli in golang to run app"
	addCommandCli(app)
	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}

func addCommandCli(app *cli.App) {
	app.Commands = []*cli.Command{
		{
			Name:   "createDb",
			Usage:  "Run command to create database",
			Action: createDb,
		},
		{
			Name:   "startApp",
			Usage:  "Run command to running app",
			Action: startApp,
		},
	}
}