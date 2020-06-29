package main

import (
	_ "log"
	"os"

	"github.com/nntruong02069999/example4/database"
	"github.com/urfave/cli/v2"
)

func createCliGolang(db *database.Db) error {
	app := cli.NewApp()
	app.Name = "cli-golang"
	app.Version = "0.0.1"
	app.Usage = "Using cli in golang to run app"
	app.Commands = []*cli.Command{
		{
			Name : "createDb",
			Usage: "Run command to create database",
			Action: func(c *cli.Context) error {
				db.ConnectDb()
				db.InitDatabase()
				return nil
			},
		},
		{
			Name : "startApp",
			Usage: "Run command to running app",
			Action : func(c *cli.Context) error {
				db.ConnectDb()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}