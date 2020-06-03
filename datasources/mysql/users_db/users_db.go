package users_db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/nishant01/mybookstore_utils-go/logger"
	"log"
)

const (
	mysqlUsersUsername = "root"
	mysqlUsersPassword = "root"
	mysqlUsersHost     = "127.0.0.1:3306"
	mysqlUsersSchema   = "users_db"
)

var (
	Client *sql.DB

	username = mysqlUsersUsername // os.Getenv(mysqlUsersUsername)
	password = mysqlUsersPassword // os.Getenv(mysqlUsersPassword)
	host     = mysqlUsersHost // os.Getenv(mysqlUsersHost)
	schema   = mysqlUsersSchema // os.Getenv(mysqlUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
