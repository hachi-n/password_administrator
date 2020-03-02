package model

import (
	"database/sql"
	"fmt"
	"github.com/hachi-n/passwd_gen/lib/config"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DBConnection *sql.DB

func init() {
	var err error
	DBConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln("database connection error.")
	}
}

func createPasswordTable() {
	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    		service_id INTERGER NOT NULL,
    		category_id INTERGER NOT NULL,
            password VARCHAR(255) NOT NULL
		)`, passwordTableName)
	_, err := DBConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}

func createServiceTable() {
	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
            name VARCHAR(255) NOT NULL
		)`, serviceTableName)
	_, err := DBConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}

func createCategoryTable() {
	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
            name VARCHAR(255) NOT NULL
		)`, categoryTableName)
	_, err := DBConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
}
