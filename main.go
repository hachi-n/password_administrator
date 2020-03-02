package main

import (
	"fmt"
	"github.com/hachi-n/passwd_gen/app/model"
	"github.com/hachi-n/passwd_gen/cli"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func init() {
	cli.Load()
}

func main() {
	if cli.Option.List {
		passwords := model.SearchPasswords()
		for _, password := range passwords {
			fmt.Println(password)
		}

		os.Exit(0)
	}

	password := model.NewPassword(cli.Option.ServiceName, cli.Option.CategoryName)
	password.Save()
	fmt.Println(password)
}

