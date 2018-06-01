package data

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaosongfu/url/config"
)

// database
var Db *sql.DB

// init Db connection
func init() {
	log.Println("database connecting ...")

	database := config.Conf.Database[config.Conf.Env]
	var err error
	//Db, err = sql.Open("mysql", "root:url@(192.168.160.3:33306)/url?charset=utf8")
	Db, err = sql.Open(database.DriverName, database.UserName+":"+database.Password+"@("+database.Host+":"+database.Port+")/"+database.Database+"?charset=utf8")
	if err != nil {
		panic(err)
	}

	// use Ping() to test db connect
	if err := Db.Ping(); err != nil {
		panic(err)
	}

	log.Println("database connect success!")
}
