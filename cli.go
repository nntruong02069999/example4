package main

import (
	_ "log"
	"os"
	"github.com/urfave/cli/v2"
)

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
