package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_USERNAME = "root"
	MYSQL_PASSWORD = ""
	MYSQL_HOST     = "127.0.0.1:3306"
	MYSQL_PORT     = "3306"
	MYSQL_SCHEMA   = "users_db"
)

var (
	Client   *sql.DB
	username = MYSQL_USERNAME // os.Getenv(MYSQL_USERS_USERNAME)
	password = MYSQL_PASSWORD // os.Getenv(MYSQL_USERS_PASSWORD)
	host     = MYSQL_HOST     // os.Getenv(MYSQL_USERS_HOST)
	port     = MYSQL_PORT
	schema   = MYSQL_SCHEMA // os.Getenv(MYSQL_USERS_SCHEMA)
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully connected")
}
