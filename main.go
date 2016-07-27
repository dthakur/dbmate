package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	loadDotEnv()

	app := NewApp()
	err := app.Run(os.Args)
	checkErr(err)
}

// NewApp creates a new command line app
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "dbmate"
	app.Usage = "A lightweight, framework-independent database migration tool."
	app.Version = Version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "migrations-dir, d",
			Value: "./db/migrations",
			Usage: "specify the directory containing migration files",
		},
		cli.StringFlag{
			Name:  "env, e",
			Value: "DATABASE_URL",
			Usage: "specify an environment variable containing the database URL",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "new",
			Usage: "Generate a new migration file",
			Action: func(ctx *cli.Context) {
				runCommand(NewCommand, ctx)
			},
		},
		{
			Name:  "up",
			Usage: "Create database (if necessary) and migrate to the latest version",
			Action: func(ctx *cli.Context) {
				runCommand(UpCommand, ctx)
			},
		},
		{
			Name:  "create",
			Usage: "Create database",
			Action: func(ctx *cli.Context) {
				runCommand(CreateCommand, ctx)
			},
		},
		{
			Name:  "drop",
			Usage: "Drop database (if it exists)",
			Action: func(ctx *cli.Context) {
				runCommand(DropCommand, ctx)
			},
		},
		{
			Name:  "migrate",
			Usage: "Migrate to the latest version",
			Action: func(ctx *cli.Context) {
				runCommand(MigrateCommand, ctx)
			},
		},
		{
			Name:  "rollback",
			Aliases: []string{"down"},
			Usage: "Rollback the most recent migration",
			Action: func(ctx *cli.Context) {
				runCommand(RollbackCommand, ctx)
			},
		},
	}

	return app
}

type command func(*cli.Context) error

func runCommand(cmd command, ctx *cli.Context) {
	err := cmd(ctx)
	checkErr(err)
}

func loadDotEnv() {
	if _, err := os.Stat(".env"); err != nil {
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
